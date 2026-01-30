package hr

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Terracode-Dev/North-Star-Server/internal/database"
	"github.com/labstack/echo/v4"
)

func (s *HRService) CheckValideteEmp(c echo.Context) error {
	var req CheckValideteEmpReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"err": err.Error(),
			"msg": "bind error",
		})
	}

	empId, err := s.q.GetEmployeeIdByEmail(c.Request().Context(), req.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"err": err.Error(),
			"msg": "invalid email",
		})
	}
	log.Printf("empId: %d", empId)

	leaveData, err := s.q.GetEmployeeLeaveBenefits(c.Request().Context(), empId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"err": err.Error(),
			"msg": "invalid leave Data",
		})
	}

	approvedLeaveCount, err := s.q.GetEmployeeApprovedLeaveCount(c.Request().Context(), database.GetEmployeeApprovedLeaveCountParams{
		EmpID:     empId,
		LeaveType: leaveData.LeaveType.String,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"err": err.Error(),
			"msg": "failed to get approved leave count",
		})
	}

	remainingLeaves := int64(leaveData.LeaveCount.Int32) - approvedLeaveCount

	res := CheckValideteEmpRes{
		EmpId:           empId,
		LeaveType:       leaveData.LeaveType.String,
		LeaveCount:      int64(leaveData.LeaveCount.Int32),
		RemainingLeaves: remainingLeaves,
	}
	return c.JSON(http.StatusOK, res)
}

// =====================================================
// CREATE LEAVE HANDLER
// =====================================================

func (s *HRService) CreateLeaveHandler(c echo.Context) error {
	var req []CreateLeaveRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if len(req) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "No leave requests provided"})
	}

	var createdLeaves []map[string]interface{}

	// Process each leave request
	for i, r := range req {
		// Parse date
		leaveDate, err := time.Parse("2006-01-02", r.LeaveDate)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": fmt.Sprintf("Invalid date format for request %d. Use YYYY-MM-DD", i+1),
			})
		}

		leaveData, err := s.q.GetEmployeeLeaveBenefits(c.Request().Context(), r.EmpID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"err": err.Error(),
				"msg": fmt.Sprintf("Invalid leave data for request %d", i+1),
			})
		}

		if leaveData.LeaveType.String == r.LeaveType {
			var maldivianTZ *time.Location
			maldivianTZ, err = time.LoadLocation("Indian/Maldives")
			if err != nil {
				maldivianTZ = time.FixedZone("MVT", 5*60*60) // UTC+5
			}

			now := time.Now().In(maldivianTZ)
			startOfYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, maldivianTZ)
			endOfYear := time.Date(now.Year(), 12, 31, 23, 59, 59, 999999999, maldivianTZ)

			params := database.GetEmployeeLeavesCountParams{
				EmpID:       r.EmpID,
				Column2:     "",                  // Empty string to ignore leave_type filter
				CONCAT:      leaveData.LeaveType, // Same value as Column2
				Column4:     startOfYear,         // Start date filter (not NULL)
				LeaveDate:   startOfYear,         // Actual start date
				Column6:     endOfYear,           // End date filter (not NULL)
				LeaveDate_2: endOfYear,           // Actual end date
			}

			dbCount, err := s.q.GetEmployeeLeavesCount(c.Request().Context(), params)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"err": err.Error(),
					"msg": fmt.Sprintf("Error getting leave count for request %d", i+1),
				})
			}

			if int64(leaveData.LeaveCount.Int32) <= dbCount {
				return c.JSON(http.StatusConflict, map[string]string{
					"err": "Leave quota exceeded",
					"msg": fmt.Sprintf("Leave count exceeded for employee %d in request %d", r.EmpID, i+1),
				})
			}
		}

		// Prepare parameters
		params := database.CreateLeaveParams{
			EmpID:     r.EmpID,
			LeaveType: r.LeaveType,
			LeaveDate: leaveDate,
			Reason:    r.Reason,
		}

		if r.AddedBy != nil {
			params.AddedBy = sql.NullInt64{Int64: *r.AddedBy, Valid: true}
		}

		result, err := s.q.CreateLeave(c.Request().Context(), params)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
				"msg":   fmt.Sprintf("Failed to create leave for request %d", i+1),
			})
		}

		id, _ := result.LastInsertId()
		createdLeaves = append(createdLeaves, map[string]interface{}{
			"employee_id": r.EmpID,
			"leave_id":    id,
			"leave_date":  r.LeaveDate,
			"leave_type":  r.LeaveType,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message":       "Leaves created successfully",
		"created_count": len(createdLeaves),
		"leaves":        createdLeaves,
	})
}

// =====================================================
// GET ALL LEAVES HANDLER
// =====================================================

func (s *HRService) GetAllLeavesHandler(c echo.Context) error {
	// Parse query parameters
	searchName := c.QueryParam("search_name")
	searchEmail := c.QueryParam("search_email")
	searchLeaveType := c.QueryParam("search_leave_type")
	dateFrom := c.QueryParam("date_from")
	dateTo := c.QueryParam("date_to")
	sortBy := c.QueryParam("sort_by") // name_asc, name_desc, email_asc, etc.

	// Pagination
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit < 1 || limit > 100 {
		limit = 10
	}

	if limit > 2147483647 { // Max value of int32
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Limit value too large",
		})
	}

	offset := (page - 1) * limit

	// Parse dates
	var parsedDateFrom, parsedDateTo time.Time
	var err error

	if dateFrom != "" {
		parsedDateFrom, err = time.Parse("2006-01-02", dateFrom)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid date_from format"})
		}
	}

	if dateTo != "" {
		parsedDateTo, err = time.Parse("2006-01-02", dateTo)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid date_to format"})
		}
	}

	// Prepare parameters for GetAllLeaves
	params := database.GetAllLeavesParams{
		Column1:   searchName,
		CONCAT:    searchName,
		CONCAT_2:  searchName,
		Column4:   searchEmail,
		CONCAT_3:  searchEmail,
		Column6:   searchLeaveType,
		LeaveType: searchLeaveType,
		Column8:   nil,
		Column10:  nil,
		Column12:  sortBy,
		Column13:  sortBy,
		Column14:  sortBy,
		Column15:  sortBy,
		Column16:  sortBy,
		Column17:  sortBy,
		Column18:  sortBy,
		Column19:  sortBy,
		Limit:     int32(limit),
		Offset:    int32(offset),
	}

	if dateFrom != "" {
		params.Column8 = parsedDateFrom
		params.LeaveDate = parsedDateFrom
	}
	if dateTo != "" {
		params.Column10 = parsedDateTo
		params.LeaveDate_2 = parsedDateTo
	}

	leaves, err := s.q.GetAllLeaves(c.Request().Context(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch leaves"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": leaves,
		"pagination": map[string]interface{}{
			"page":   page,
			"limit":  limit,
			"offset": offset,
		},
	})
}

// =====================================================
// GET LEAVE BY ID HANDLER
// =====================================================

func (s *HRService) GetLeaveByIdHandler(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid leave ID"})
	}

	leave, err := s.q.GetLeaveById(c.Request().Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Leave not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch leave"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": leave,
	})
}

// =====================================================
// UPDATE LEAVE HANDLER
// =====================================================

type UpdateLeaveRequest struct {
	LeaveType string `json:"leave_type" validate:"required"`
	LeaveDate string `json:"leave_date" validate:"required"`
	Reason    string `json:"reason" validate:"required"`
	AddedBy   *int64 `json:"added_by,omitempty"`
}

func (s *HRService) UpdateLeaveHandler(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid leave ID"})
	}

	var req UpdateLeaveRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Parse date
	leaveDate, err := time.Parse("2006-01-02", req.LeaveDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid date format. Use YYYY-MM-DD"})
	}

	// Prepare parameters
	params := database.UpdateLeaveParams{
		ID:        id,
		LeaveType: req.LeaveType,
		LeaveDate: leaveDate,
		Reason:    req.Reason,
	}

	if req.AddedBy != nil {
		params.AddedBy = sql.NullInt64{Int64: *req.AddedBy, Valid: true}
	}

	err = s.q.UpdateLeave(c.Request().Context(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update leave"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Leave updated successfully",
	})
}

// =====================================================
// DELETE LEAVE HANDLER
// =====================================================

func (s *HRService) DeleteLeaveHandler(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid leave ID"})
	}

	err = s.q.DeleteLeave(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete leave"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Leave deleted successfully",
	})
}

// =====================================================
// GET EMPLOYEE LEAVES HANDLER
// =====================================================

func (s *HRService) GetEmployeeLeavesHandler(c echo.Context) error {
	empId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid employee ID"})
	}

	// Parse query parameters
	searchLeaveType := c.QueryParam("search_leave_type")
	year := c.QueryParam("year")
	sortBy := c.QueryParam("sort_by") // date_asc, date_desc, type_asc, type_desc

	// Pagination
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit < 1 || limit > 100 {
		limit = 10
	}

	if limit > 2147483647 { // Max value of int32
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Limit value too large",
		})
	}

	offset := (page - 1) * limit

	// Parse year
	var parsedYear time.Time
	if year != "" {
		parsedYear, err = time.Parse("2006", year)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid year format"})
		}
	}

	// Prepare parameters
	params := database.GetEmployeeLeavesParams{
		EmpID:   empId,
		Column2: searchLeaveType,
		CONCAT:  searchLeaveType,
		Column4: nil,
		Column6: sortBy,
		Column7: sortBy,
		Column8: sortBy,
		Column9: sortBy,
		Limit:   int32(limit),
		Offset:  int32(offset),
	}

	if year != "" {
		params.Column4 = parsedYear
		params.LeaveDate = parsedYear
	}

	leaves, err := s.q.GetEmployeeLeaves(c.Request().Context(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch employee leaves"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": leaves,
		"pagination": map[string]interface{}{
			"page":   page,
			"limit":  limit,
			"offset": offset,
		},
	})
}

// =====================================================
// GET EMPLOYEE LEAVE BENEFITS HANDLER
// =====================================================

func (s *HRService) GetEmployeeLeaveBenefitsHandler(c echo.Context) error {
	empId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid employee ID"})
	}

	benefits, err := s.q.GetEmployeeLeaveBenefits(c.Request().Context(), empId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch employee leave benefits"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": benefits,
	})
}

func (s *HRService) GetLeavesHandlerEMP(c echo.Context) error {
	empId, ok := getInt64(c.Get("user_id"))
	if !ok {
		log.Print("error parsing user_id")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid employee ID"})
	}

	// Parse query parameters
	searchLeaveType := c.QueryParam("search_leave_type")
	year := c.QueryParam("year")
	sortBy := c.QueryParam("sort_by") // date_asc, date_desc, type_asc, type_desc

	// Pagination
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit < 1 || limit > 100 {
		limit = 10
	}
	if limit > 2147483647 { // Max value of int32
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Limit value too large",
		})
	}
	offset := (page - 1) * limit

	// Parse year
	var yearPtr *time.Time
	if year != "" {
		parsedYear, err := time.Parse("2006", year)
		if err != nil {
			log.Printf("Invalid year format: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid year format"})
		}
		yearPtr = &parsedYear
	}

	// Prepare parameters
	params := database.GetEmployeeLeavesEMPParams{
		EmpID:   empId,
		Column2: searchLeaveType,
		CONCAT:  searchLeaveType,
		Column4: yearPtr,
		Column6: sortBy,
		Column7: sortBy,
		Column8: sortBy,
		Column9: sortBy,
		Limit:   int32(limit),
		Offset:  int32(offset),
	}

	// Only set LeaveDate if year is provided
	if yearPtr != nil {
		params.LeaveDate = *yearPtr
	}

	// Fetch leaves from database
	leaves, err := s.q.GetEmployeeLeavesEMP(c.Request().Context(), params)
	if err != nil {
		log.Printf("Failed to fetch employee leaves for empId %d: %v, params: %+v", empId, err, params)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch employee leaves",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": leaves,
		"pagination": map[string]interface{}{
			"page":   page,
			"limit":  limit,
			"offset": offset,
		},
	})
}

func (s *HRService) CreateLeaveHandlerEMP(c echo.Context) error {
	empId, ok := getInt64(c.Get("user_id"))
	if !ok {
		log.Print("error parsing user_id")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid employee ID"})
	}

	var req CreateLeaveRequestEMP
	if err := c.Bind(&req); err != nil {
		log.Printf("Error binding request: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	leaveDate, err := time.Parse("2006-01-02", req.LeaveDate)
	if err != nil {
		log.Printf("Error parsing leave date: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid date format for request. Use YYYY-MM-DD",
		})
	}

	// Get employee leave benefits
	leaveData, err := s.q.GetEmployeeLeaveBenefits(c.Request().Context(), empId)
	if err != nil {
		// Handle case where employee has no leave benefits assigned
		if err == sql.ErrNoRows {
			log.Printf("No leave benefits found for empId %d", empId)
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "No leave type assigned to this employee. Please contact HR to set up your leave benefits.",
			})
		}
		// Handle other database errors
		log.Printf("Error getting leave benefits for empId %d: %v", empId, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch employee leave benefits",
		})
	}

	// Validate leave type exists
	if !leaveData.LeaveType.Valid || leaveData.LeaveType.String == "" {
		log.Printf("No valid leave type found for empId %d", empId)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "No leave type assigned to employee. Please contact HR.",
		})
	}

	// Validate leave count exists
	if !leaveData.LeaveCount.Valid {
		log.Printf("No valid leave count found for empId %d", empId)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "No leave quota assigned to employee. Please contact HR.",
		})
	}

	// Setup timezone
	maldivianTZ, err := time.LoadLocation("Indian/Maldives")
	if err != nil {
		log.Printf("Error loading Maldives timezone, using UTC+5: %v", err)
		maldivianTZ = time.FixedZone("MVT", 5*60*60) // UTC+5
	}

	now := time.Now().In(maldivianTZ)
	startOfYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, maldivianTZ)
	endOfYear := time.Date(now.Year(), 12, 31, 23, 59, 59, 999999999, maldivianTZ)

	// Prepare count parameters
	params := database.GetEmployeeLeavesCountParams{
		EmpID:       empId,
		Column2:     leaveData.LeaveType.String,
		CONCAT:      leaveData.LeaveType.String,
		Column4:     startOfYear,
		LeaveDate:   startOfYear,
		Column6:     endOfYear,
		LeaveDate_2: endOfYear,
	}

	// Get current leave count for the year
	dbCount, err := s.q.GetEmployeeLeavesCount(c.Request().Context(), params)
	if err != nil {
		log.Printf("Error getting leave count for empId %d: %v, params: %+v", empId, err, params)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch employee leave count",
		})
	}

	// Check if quota exceeded
	if int64(leaveData.LeaveCount.Int32) <= dbCount {
		log.Printf("Leave quota exceeded for empId %d: used=%d, total=%d", empId, dbCount, leaveData.LeaveCount.Int32)
		return c.JSON(http.StatusConflict, map[string]string{
			"error": fmt.Sprintf("Leave quota exceeded. You have used %d out of %d leaves this year.", dbCount, leaveData.LeaveCount.Int32),
		})
	}

	// Prepare insert parameters
	dbParam := database.CreateLeaveParams{
		EmpID:     empId,
		LeaveType: leaveData.LeaveType.String,
		LeaveDate: leaveDate,
		Reason:    req.Reason,
		AddedBy: sql.NullInt64{
			Valid: true,
			Int64: empId,
		},
	}

	// Create leave
	result, err := s.q.CreateLeave(c.Request().Context(), dbParam)
	if err != nil {
		log.Printf("Error creating leave for empId %d: %v, params: %+v", empId, err, dbParam)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create leave",
		})
	}

	// Get the inserted ID
	leaveId, _ := result.LastInsertId()
	log.Printf("Leave created successfully for empId %d, leaveId: %d", empId, leaveId)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message":  "Leave created successfully",
		"leave_id": leaveId,
	})
}

func (s *HRService) UpdateLeaveHandlerEMP(c echo.Context) error {
	id, ok := getInt64(c.Get("user_id"))
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid employee ID"})
	}
	var req UpdateLeaveRequestEMP
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Parse date
	leaveDate, err := time.Parse("2006-01-02", req.LeaveDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid date format. Use YYYY-MM-DD"})
	}

	// Prepare parameters
	params := database.UpdateLeaveEMPParams{
		ID:        id,
		LeaveDate: leaveDate,
		Reason:    req.Reason,
	}

	err = s.q.UpdateLeaveEMP(c.Request().Context(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update leave"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Leave updated successfully",
	})
}

func (s *HRService) DeleteLeaveHandlerEMP(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid leave ID"})
	}

	err = s.q.DeleteLeave(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete leave"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Leave deleted successfully",
	})
}
