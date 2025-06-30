package hr

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// Request/Response structs
type CreateLeaveRequest struct {
	EmpID     int64  `json:"emp_id" validate:"required"`
	LeaveType string `json:"leave_type" validate:"required"`
	LeaveDate string `json:"leave_date" validate:"required"` // YYYY-MM-DD format
	Reason    string `json:"reason" validate:"required"`
	AddedBy   *int64 `json:"added_by"`
}

type UpdateLeaveRequest struct {
	LeaveType string `json:"leave_type" validate:"required"`
	LeaveDate string `json:"leave_date" validate:"required"`
	Reason    string `json:"reason" validate:"required"`
}

type LeaveResponse struct {
	ID         int64      `json:"id"`
	EmpID      int64      `json:"emp_id"`
	LeaveType  string     `json:"leave_type"`
	LeaveDate  time.Time  `json:"leave_date"`
	Reason     string     `json:"reason"`
	CreateDate *time.Time `json:"create_date,omitempty"`
	AddedBy    *int64     `json:"added_by,omitempty"`
}

type PaginatedLeavesResponse struct {
	Data       []LeaveResponse `json:"data"`
	Total      int64           `json:"total"`
	Page       int             `json:"page"`
	Limit      int             `json:"limit"`
	TotalPages int             `json:"total_pages"`
}

type LeaveSummaryResponse struct {
	EmpID          int64  `json:"emp_id"`
	LeaveType      string `json:"leave_type"`
	UsedCount      int64  `json:"used_count"`
	AllowedCount   int32  `json:"allowed_count"`
	RemainingCount int32  `json:"remaining_count"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
}

type LeaveValidationResponse struct {
	CanCreateLeave bool  `json:"can_create_leave"`
	TotalAllowed   int32 `json:"total_allowed"`
	AlreadyUsed    int64 `json:"already_used"`
	Remaining      int64 `json:"remaining"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// Helper function to send error response
func (s *HRService) sendError(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
	})
}

// Helper function to convert interface{} to int64
// func interfaceToInt64(val interface{}) int64 {
// 	switch v := val.(type) {
// 	case int64:
// 		return v
// 	case int32:
// 		return int64(v)
// 	case int:
// 		return int64(v)
// 	case float64:
// 		return int64(v)
// 	default:
// 		return 0
// 	}
// }
