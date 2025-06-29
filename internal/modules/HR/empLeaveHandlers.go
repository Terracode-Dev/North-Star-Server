package hr

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/Terracode-Dev/North-Star-Server/internal/database"
	"github.com/labstack/echo/v4"
)

// 1. Create Leave
func (s *HRService) CreateLeave(c echo.Context) error {
	var req CreateLeaveRequest
	if err := c.Bind(&req); err != nil {
		return s.sendError(c, http.StatusBadRequest, "Invalid request body")
	}

	// Validate request
	if req.EmpID == 0 || req.LeaveType == "" || req.LeaveDate == "" || req.Reason == "" {
		return s.sendError(c, http.StatusBadRequest, "Missing required fields")
	}

	// Parse date
	leaveDate, err := parseDate(req.LeaveDate)
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, "Invalid date format. Use YYYY-MM-DD")
	}

	// Check leave count for validation
	validation, err := s.q.CheckLeaveCountForYear(c.Request().Context(), database.CheckLeaveCountForYearParams{
		EmployeeID: req.EmpID,
		LeaveType:  req.LeaveType,
	})
	if err != nil {
		return s.sendError(c, http.StatusInternalServerError, "Failed to validate leave creation")
	}

	usedLeaves := interfaceToInt64(validation.UsedLeaves)
	if usedLeaves >= int64(validation.TotalAllowed) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Leave limit exceeded",
			"validation": LeaveValidationResponse{
				CanCreateLeave: false,
				TotalAllowed:   validation.TotalAllowed,
				AlreadyUsed:    usedLeaves,
				Remaining:      int64(validation.TotalAllowed) - usedLeaves,
			},
		})
	}

	// Create leave
	var addedBy sql.NullInt64
	if req.AddedBy != nil {
		addedBy = sql.NullInt64{Int64: *req.AddedBy, Valid: true}
	}

	err = s.q.CreateLeave(c.Request().Context(), database.CreateLeaveParams{
		EmpID:     req.EmpID,
		LeaveType: req.LeaveType,
		LeaveDate: leaveDate,
		Reason:    req.Reason,
		AddedBy:   addedBy,
	})
	if err != nil {
		return s.sendError(c, http.StatusInternalServerError, "Failed to create leave")
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Leave created successfully",
	})
}

// 2. Update Leave
func (s *HRService) UpdateLeave(c echo.Context) error {
	leaveID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, "Invalid leave ID")
	}

	empID, err := strconv.ParseInt(c.Param("emp_id"), 10, 64)
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, "Invalid employee ID")
	}

	var req UpdateLeaveRequest
	if err := c.Bind(&req); err != nil {
		return s.sendError(c, http.StatusBadRequest, "Invalid request body")
	}

	// Parse date
	leaveDate, err := parseDate(req.LeaveDate)
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, "Invalid date format. Use YYYY-MM-DD")
	}

	// Update leave
	err = s.q.UpdateLeave(c.Request().Context(), database.UpdateLeaveParams{
		LeaveType: req.LeaveType,
		LeaveDate: leaveDate,
		Reason:    req.Reason,
		ID:        leaveID,
		EmpID:     empID,
	})
	if err != nil {
		return s.sendError(c, http.StatusInternalServerError, "Failed to update leave")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Leave updated successfully",
	})
}

// 3. Delete Leave
func (s *HRService) DeleteLeave(c echo.Context) error {
	leaveID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, "Invalid leave ID")
	}

	empID, err := strconv.ParseInt(c.Param("emp_id"), 10, 64)
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, "Invalid employee ID")
	}

	err = s.q.DeleteLeave(c.Request().Context(), database.DeleteLeaveParams{
		ID:    leaveID,
		EmpID: empID,
	})
	if err != nil {
		return s.sendError(c, http.StatusInternalServerError, "Failed to delete leave")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Leave deleted successfully",
	})
}

// 4. Get Employee Leaves
func (s *HRService) GetEmployeeLeaves(c echo.Context) error {
	empID, err := strconv.ParseInt(c.Param("emp_id"), 10, 64)
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, "Invalid employee ID")
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	leaves, err := s.q.GetEmployeeLeaves(c.Request().Context(), database.GetEmployeeLeavesParams{
		EmpID:  empID,
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return s.sendError(c, http.StatusInternalServerError, "Failed to get employee leaves")
	}

	var leaveResponses []LeaveResponse
	var total int64
	for _, leave := range leaves {
		total = interfaceToInt64(leave.TotalCount)

		var createDate *time.Time
		if leave.CreateDate.Valid {
			createDate = &leave.CreateDate.Time
		}

		leaveResponses = append(leaveResponses, LeaveResponse{
			ID:         leave.ID,
			EmpID:      leave.EmpID,
			LeaveType:  leave.LeaveType,
			LeaveDate:  leave.LeaveDate,
			Reason:     leave.Reason,
			CreateDate: createDate,
		})
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	response := PaginatedLeavesResponse{
		Data:       leaveResponses,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}

	return c.JSON(http.StatusOK, response)
}

// 5. Get Leave Summary
func (s *HRService) GetLeaveSummary(c echo.Context) error {
	empID, _ := strconv.ParseInt(c.QueryParam("emp_id"), 10, 64)

	var column1 interface{} = 0
	if empID > 0 {
		column1 = empID
	}

	summary, err := s.q.GetLeaveSummaryByEmployee(c.Request().Context(), database.GetLeaveSummaryByEmployeeParams{
		Column1: column1,
		EmpID:   empID,
	})
	if err != nil {
		return s.sendError(c, http.StatusInternalServerError, "Failed to get leave summary")
	}

	var summaryResponses []LeaveSummaryResponse
	for _, item := range summary {
		summaryResponses = append(summaryResponses, LeaveSummaryResponse{
			EmpID:          item.EmpID,
			LeaveType:      item.LeaveType,
			UsedCount:      item.UsedCount,
			AllowedCount:   item.AllowedCount,
			RemainingCount: item.RemainingCount,
			FirstName:      item.FirstName,
			LastName:       item.LastName,
		})
	}

	return c.JSON(http.StatusOK, summaryResponses)
}

// 6. Get Leave Types for Employee
func (s *HRService) GetLeaveTypesForEmployee(c echo.Context) error {
	empID, err := strconv.ParseInt(c.Param("emp_id"), 10, 64)
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, "Invalid employee ID")
	}

	leaveTypes, err := s.q.GetLeaveTypesForEmployee(c.Request().Context(), empID)
	if err != nil {
		return s.sendError(c, http.StatusInternalServerError, "Failed to get leave types")
	}

	return c.JSON(http.StatusOK, leaveTypes)
}

// 7. Get Employee Leave Benefits
func (s *HRService) GetEmployeeLeaveBenefits(c echo.Context) error {
	empID, err := strconv.ParseInt(c.Param("emp_id"), 10, 64)
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, "Invalid employee ID")
	}

	benefits, err := s.q.GetEmployeeLeaveBenefits(c.Request().Context(), empID)
	if err != nil {
		return s.sendError(c, http.StatusInternalServerError, "Failed to get employee leave benefits")
	}

	return c.JSON(http.StatusOK, benefits)
}

// 8. Validate Leave Creation
func (s *HRService) ValidateLeaveCreation(c echo.Context) error {
	empID, err := strconv.ParseInt(c.QueryParam("emp_id"), 10, 64)
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, "Invalid employee ID")
	}

	leaveType := c.QueryParam("leave_type")
	if leaveType == "" {
		return s.sendError(c, http.StatusBadRequest, "Leave type is required")
	}

	validation, err := s.q.CheckLeaveCountForYear(c.Request().Context(), database.CheckLeaveCountForYearParams{
		EmployeeID: empID,
		LeaveType:  leaveType,
	})
	if err != nil {
		return s.sendError(c, http.StatusInternalServerError, "Failed to validate leave creation")
	}

	usedLeaves := interfaceToInt64(validation.UsedLeaves)
	remaining := int64(validation.TotalAllowed) - usedLeaves

	response := LeaveValidationResponse{
		CanCreateLeave: remaining > 0,
		TotalAllowed:   validation.TotalAllowed,
		AlreadyUsed:    usedLeaves,
		Remaining:      remaining,
	}

	return c.JSON(http.StatusOK, response)
}
