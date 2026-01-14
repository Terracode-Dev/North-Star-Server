package hr

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/Terracode-Dev/North-Star-Server/internal/database"
	"github.com/labstack/echo/v4"
)

// Request/Response DTOs
type CreateAirTicketReqDTO struct {
	PassengerName  string `json:"passenger_name"`
	PassengerEmail string `json:"passenger_email"`
	PassportNumber string `json:"passport_number"`
	DepartureDate  string `json:"departure_date"` // YYYY-MM-DD
	ReturnDate     string `json:"return_date"`    // YYYY-MM-DD
	DepartureCity  string `json:"departure_city"`
	ArrivalCity    string `json:"arrival_city"`
	Reason         string `json:"reason"`
}

type UpdateAirTicketReqDTO struct {
	PassengerName  string `json:"passenger_name"`
	PassengerEmail string `json:"passenger_email"`
	PassportNumber string `json:"passport_number"`
	DepartureDate  string `json:"departure_date"`
	ReturnDate     string `json:"return_date"`
	DepartureCity  string `json:"departure_city"`
	ArrivalCity    string `json:"arrival_city"`
	Reason         string `json:"reason"`
}

type UpdateStatusDTO struct {
	Status string `json:"status"`
}

type PaginationResponse struct {
	Data       interface{} `json:"data"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
}

// CreateAirTicketRequest - Employee creates their own air ticket request
func (h *HRService) CreateAirTicketRequest(c echo.Context) error {
	userID := c.Get("user_id").(int64)
	branchID := c.Get("branch").(int64)

	var req CreateAirTicketReqDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Parse dates
	departureDate, err := time.Parse("2006-01-02", req.DepartureDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid departure date format"})
	}

	returnDate, err := time.Parse("2006-01-02", req.ReturnDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid return date format"})
	}

	// Validate dates
	if returnDate.Before(departureDate) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Return date must be after departure date"})
	}

	// Create request
	result, err := h.q.CreateEmpAirticketReq(c.Request().Context(), database.CreateEmpAirticketReqParams{
		PassengerName:  req.PassengerName,
		PassengerEmail: req.PassengerEmail,
		PassportNumber: req.PassportNumber,
		DepartureDate:  departureDate,
		ReturnDate:     returnDate,
		DepartureCity:  req.DepartureCity,
		ArrivalCity:    req.ArrivalCity,
		Reason:         req.Reason,
		EmpID:          userID,
		BranchID:       branchID,
		Status:         database.EmpAirticketReqStatusPending,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create request"})
	}

	id, _ := result.LastInsertId()
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Air ticket request created successfully",
		"id":      id,
	})
}

// GetMyAirTicketRequests - Employee gets their own requests (paginated)
func (h *HRService) GetMyAirTicketRequests(c echo.Context) error {
	userID := c.Get("user_id").(int64)
	branchID := c.Get("branch").(int64)

	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	// Get total count
	total, err := h.q.CountEmpAirticketReqByEmpAndBranch(c.Request().Context(), database.CountEmpAirticketReqByEmpAndBranchParams{
		EmpID:    userID,
		BranchID: branchID,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch count"})
	}

	// Get requests
	requests, err := h.q.GetEmpAirticketReqByEmpAndBranch(c.Request().Context(), database.GetEmpAirticketReqByEmpAndBranchParams{
		EmpID:    userID,
		BranchID: branchID,
		Limit:    int32(pageSize),
		Offset:   int32(offset),
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch requests"})
	}

	totalPages := int(total) / pageSize
	if int(total)%pageSize != 0 {
		totalPages++
	}

	return c.JSON(http.StatusOK, PaginationResponse{
		Data:       requests,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	})
}

// GetBranchAirTicketRequests - Admin/Manager gets all requests for their branch (paginated)
func (h *HRService) GetBranchAirTicketRequests(c echo.Context) error {
	branchID := c.Get("branch").(int64)
	role := c.Get("role").(string)

	// Only admins/managers can view all branch requests
	if role != "admin" && role != "manager" {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "Insufficient permissions"})
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	// Get total count
	total, err := h.q.CountEmpAirticketReqByBranch(c.Request().Context(), branchID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch count"})
	}

	// Get requests
	requests, err := h.q.GetEmpAirticketReqByBranch(c.Request().Context(), database.GetEmpAirticketReqByBranchParams{
		BranchID: branchID,
		Limit:    int32(pageSize),
		Offset:   int32(offset),
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch requests"})
	}

	totalPages := int(total) / pageSize
	if int(total)%pageSize != 0 {
		totalPages++
	}

	return c.JSON(http.StatusOK, PaginationResponse{
		Data:       requests,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	})
}

// GetAirTicketRequestByID - Get single request (with authorization check)
func (h *HRService) GetAirTicketRequestByID(c echo.Context) error {
	userID := c.Get("user_id").(int64)
	role := c.Get("role").(string)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	request, err := h.q.GetEmpAirticketReqByID(c.Request().Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Request not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch request"})
	}

	// Authorization: Only owner or admin/manager can view
	if request.EmpID != userID && role != "admin" && role != "manager" {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "Access denied"})
	}

	return c.JSON(http.StatusOK, request)
}

// UpdateAirTicketRequest - Employee updates their own request (only if pending)
func (h *HRService) UpdateAirTicketRequest(c echo.Context) error {
	userID := c.Get("user_id").(int64)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var req UpdateAirTicketReqDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Check if request exists and belongs to user
	existing, err := h.q.GetEmpAirticketReqByID(c.Request().Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Request not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch request"})
	}

	if existing.EmpID != userID {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "Access denied"})
	}

	// Only allow updates if status is pending
	if existing.Status != database.EmpAirticketReqStatusPending {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Cannot update request that is not pending"})
	}

	// Parse dates
	departureDate, err := time.Parse("2006-01-02", req.DepartureDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid departure date format"})
	}

	returnDate, err := time.Parse("2006-01-02", req.ReturnDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid return date format"})
	}

	if returnDate.Before(departureDate) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Return date must be after departure date"})
	}

	// Update request
	err = h.q.UpdateEmpAirticketReq(c.Request().Context(), database.UpdateEmpAirticketReqParams{
		PassengerName:  req.PassengerName,
		PassengerEmail: req.PassengerEmail,
		PassportNumber: req.PassportNumber,
		DepartureDate:  departureDate,
		ReturnDate:     returnDate,
		DepartureCity:  req.DepartureCity,
		ArrivalCity:    req.ArrivalCity,
		Reason:         req.Reason,
		EmpID:          existing.EmpID,
		BranchID:       existing.BranchID,
		ID:             id,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update request"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Request updated successfully"})
}

// UpdateAirTicketRequestStatus - Admin/Manager updates status
func (h *HRService) UpdateAirTicketRequestStatus(c echo.Context) error {
	role := c.Get("role").(string)

	// Only admins/managers can update status
	if role != "admin" && role != "manager" {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "Insufficient permissions"})
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var req UpdateStatusDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Check if request exists
	_, err = h.q.GetEmpAirticketReqByID(c.Request().Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Request not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch request"})
	}

	// Convert string to enum
	var status database.EmpAirticketReqStatus
	switch req.Status {
	case "pending":
		status = database.EmpAirticketReqStatusPending
	case "approved":
		status = database.EmpAirticketReqStatusApproved
	case "rejected":
		status = database.EmpAirticketReqStatusRejected
	default:
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid status"})
	}

	err = h.q.SetEmpAirticketReqStatus(c.Request().Context(), database.SetEmpAirticketReqStatusParams{
		Status: status,
		ID:     id,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update status"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Status updated successfully"})
}

// DeleteAirTicketRequest - Admin/Manager deletes a request
func (h *HRService) DeleteAirTicketRequest(c echo.Context) error {
	role := c.Get("role").(string)

	// Only admins can delete
	if role != "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "Insufficient permissions"})
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	// Check if request exists
	_, err = h.q.GetEmpAirticketReqByID(c.Request().Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Request not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch request"})
	}

	err = h.q.DeleteEmpAirticketReq(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete request"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Request deleted successfully"})
}

// GetAirTicketRequestsByBranchAndStatus - Get by specific branch and status
func (h *HRService) GetAirTicketRequestsByBranchAndStatus(c echo.Context) error {
	branchID := c.Get("branch_id").(int64)

	status := c.QueryParam("status")
	if status == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Status is required"})
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	var statusEnum database.EmpAirticketReqStatus
	switch status {
	case "pending":
		statusEnum = database.EmpAirticketReqStatusPending
	case "approved":
		statusEnum = database.EmpAirticketReqStatusApproved
	case "rejected":
		statusEnum = database.EmpAirticketReqStatusRejected
	default:
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid status"})
	}
	var total int64
	var err error
	var requests []database.EmpAirticketReq
	if branchID == 1 {
		total, err = h.q.CountEmpAirticketReqByStatus(c.Request().Context(), statusEnum)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch count"})
		}

		requests, err = h.q.GetEmpAirticketReqByStatus(c.Request().Context(), database.GetEmpAirticketReqByStatusParams{
			Status: statusEnum,
			Limit:  int32(pageSize),
			Offset: int32(offset),
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch requests"})
		}
	} else {
		total, err = h.q.CountEmpAirticketReqByBranchAndStatus(c.Request().Context(), database.CountEmpAirticketReqByBranchAndStatusParams{
			BranchID: branchID,
			Status:   statusEnum,
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch count"})
		}

		requests, err = h.q.GetEmpAirticketReqByBranchAndStatus(c.Request().Context(), database.GetEmpAirticketReqByBranchAndStatusParams{
			BranchID: branchID,
			Status:   statusEnum,
			Limit:    int32(pageSize),
			Offset:   int32(offset),
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch requests"})
		}
	}

	totalPages := int(total) / pageSize
	if int(total)%pageSize != 0 {
		totalPages++
	}

	return c.JSON(http.StatusOK, PaginationResponse{
		Data:       requests,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	})
}
