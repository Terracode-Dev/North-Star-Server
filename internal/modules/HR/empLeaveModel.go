package hr

type CreateLeaveRequest struct {
	EmpID     int64  `json:"emp_id" validate:"required"`
	LeaveType string `json:"leave_type" validate:"required"`
	LeaveDate string `json:"leave_date" validate:"required"` // "2025-07-15"
	Reason    string `json:"reason" validate:"required"`
	AddedBy   *int64 `json:"added_by,omitempty"`
}

type CreateLeaveRequestEMP struct {
	// LeaveType string `json:"leave_type" validate:"required"`
	LeaveDate string `json:"leave_date" validate:"required"` // "2025-07-15"
	Reason    string `json:"reason" validate:"required"`
}

type UpdateLeaveRequestEMP struct {
	LeaveDate string `json:"leave_date" validate:"required"`
	Reason    string `json:"reason" validate:"required"`
	AddedBy   *int64 `json:"added_by,omitempty"`
}

type CheckValideteEmpReq struct {
	Email string `json:"email"`
}

type CheckValideteEmpRes struct {
	EmpId           int64  `json:"emp_id"`
	LeaveType       string `json:"leave_type"`
	LeaveCount      int64  `json:"leave_count"`
	RemainingLeaves int64  `json:"remaining_leaves"` // New field for remaining leaves
}
