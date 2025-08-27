package hr

import (
	"database/sql"
	"time"
	"github.com/Terracode-Dev/North-Star-Server/internal/database"
)

// Request/Response Structs
type EmployeeResponse struct {
	ID         int64  `json:"id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Position   string `json:"position"`
}

type GetEmployeeByEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type CreateScheduleRequest struct {
	EmpID      int64                    `json:"emp_id" validate:"required"`
	Weekly     WeeklySchedule           `json:"weekly_schedule"`
	Additional []AdditionalScheduleItem `json:"additional_schedule,omitempty"`
}

type WeeklySchedule struct {
	Monday    *DaySchedule `json:"monday,omitempty"`
	Tuesday   *DaySchedule `json:"tuesday,omitempty"`
	Wednesday *DaySchedule `json:"wednesday,omitempty"`
	Thursday  *DaySchedule `json:"thursday,omitempty"`
	Friday    *DaySchedule `json:"friday,omitempty"`
	Saturday  *DaySchedule `json:"saturday,omitempty"`
	Sunday    *DaySchedule `json:"sunday,omitempty"`
}

type DaySchedule struct {
	IsWorking bool   `json:"is_working"`
	FromTime  string `json:"from_time,omitempty"` // Format: "09:00:00"
	ToTime    string `json:"to_time,omitempty"`   // Format: "17:00:00"
}

type AdditionalScheduleItem struct {
	Date     string `json:"date" validate:"required"` // Format: "2024-01-15"
	FromTime string `json:"from_time,omitempty"`      // Format: "09:00:00"
	ToTime   string `json:"to_time,omitempty"`        // Format: "17:00:00"
}

type UpdateScheduleRequest struct {
	Weekly     *WeeklySchedule          `json:"weekly_schedule,omitempty"`
	Additional []AdditionalScheduleItem `json:"additional_schedule,omitempty"`
}

type EmployeeListRequest struct {
	Page      int    `query:"page"`
	Limit     int    `query:"limit"`
	FirstName string `query:"first_name"`
	LastName  string `query:"last_name"`
	Email     string `query:"email"`
	SortBy    string `query:"sort_by"` // first_name, last_name, email, work_days
	Year      int    `query:"year"`
}

type EmployeeListResponse struct {
	Data       []EmployeeListItem `json:"data"`
	Pagination PaginationInfo     `json:"pagination"`
	Filters    FilterInfo         `json:"filters"`
}

type EmployeeListItem struct {
	ID              int64  `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	WorkDaysForYear int32  `json:"work_days_for_year"`
}

type PaginationInfo struct {
	CurrentPage  int   `json:"current_page"`
	PageSize     int   `json:"page_size"`
	TotalRecords int64 `json:"total_records"`
	TotalPages   int   `json:"total_pages"`
	HasNext      bool  `json:"has_next"`
	HasPrevious  bool  `json:"has_previous"`
}

type FilterInfo struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	SortBy    string `json:"sort_by,omitempty"`
	Year      int    `json:"year"`
}

type WorkDaysBreakdownResponse struct {
	ID                   int64  `json:"id"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	Email                string `json:"email"`
	WeeklyWorkDays       int64  `json:"weekly_work_days"`
	BaseYearlyDays       int64  `json:"base_yearly_days"`
	AdditionalDays       int64  `json:"additional_days"`
	TotalWorkDaysForYear int32  `json:"total_work_days_for_year"`
}

// Utility functions
func parseTime(timeStr string) (sql.NullTime, error) {
	if timeStr == "" {
		return sql.NullTime{Valid: false}, nil
	}

	t, err := time.Parse("15:04:05", timeStr)
	if err != nil {
		return sql.NullTime{Valid: false}, err
	}

	baseTime := time.Date(1970, 1, 1, t.Hour(), t.Minute(), t.Second(), 0, time.UTC)

	return sql.NullTime{Time: baseTime, Valid: true}, nil
}

func parseDate(dateStr string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func interfaceToString(val interface{}) string {
	if val == nil {
		return ""
	}
	if str, ok := val.(string); ok {
		return str
	}
	return ""
}

func interfaceToInt32(val interface{}) int32 {
	if val == nil {
		return 0
	}
	if i, ok := val.(int32); ok {
		return i
	}
	if i, ok := val.(int); ok {
		return int32(i)
	}
	return 0
}

func interfaceToInt64(val interface{}) int64 {
	if val == nil {
		return 0
	}
	if i, ok := val.(int64); ok {
		return i
	}
	if i, ok := val.(int); ok {
		return int64(i)
	}
	return 0
}

type GetEmployeeAttendanceReqParams struct {
	EmpID      int64       `json:"emp_id"`
	CreateDate string   `json:"create_date"`
	Limit      int32       `json:"limit"`
	Offset     int32       `json:"offset"`
}

func (r *GetEmployeeAttendanceReqParams) ToInsufficientDBStruct() (database.GetInsufficientAttendanceParams, error) {
	if r.CreateDate == "" {
		return database.GetInsufficientAttendanceParams{
			EmpID:      r.EmpID,
			Column2:    sql.NullTime{},
			Limit:      r.Limit,
			Offset:     r.Offset,
		}, nil
	}
	createDate, err := parseDate(r.CreateDate)
	if err != nil {
		return database.GetInsufficientAttendanceParams{}, err
	}

	return database.GetInsufficientAttendanceParams{
		EmpID:      r.EmpID,
		Column2:    createDate,
		CreateDate: createDate,
		Limit:      r.Limit,
		Offset:     r.Offset,
	}, nil
}

func (r *GetEmployeeAttendanceReqParams) ToLateDBStruct() (database.GetLateAttendanceParams, error) {
	if r.CreateDate == "" {
		return database.GetLateAttendanceParams{
			EmpID:      r.EmpID,
			Column2:    sql.NullTime{},
			Limit:      r.Limit,
			Offset:     r.Offset,
		}, nil
	}
	createDate, err := parseDate(r.CreateDate)
	if err != nil {
		return database.GetLateAttendanceParams{}, err
	}

	return database.GetLateAttendanceParams{
		EmpID:      r.EmpID,
		Column2:    createDate,
		CreateDate: createDate,
		Limit:      r.Limit,
		Offset:     r.Offset,
	}, nil
}

func (r *GetEmployeeAttendanceReqParams) ToNormalDBStruct() (database.GetNormalAttendanceParams, error) {
	if r.CreateDate == "" {
		return database.GetNormalAttendanceParams{
			EmpID:      r.EmpID,
			Column2:    sql.NullTime{},
			Limit:      r.Limit,
			Offset:     r.Offset,
		}, nil
	}
	createDate, err := parseDate(r.CreateDate)
	if err != nil {
		return database.GetNormalAttendanceParams{}, err
	}

	return database.GetNormalAttendanceParams{
		EmpID:      r.EmpID,
		Column2:    createDate,
		CreateDate: createDate,
		Limit:      r.Limit,
		Offset:     r.Offset,
	}, nil
}

func (r *GetEmployeeAttendanceReqParams) ToAllDBStruct() (database.GetAllAttendanceParams, error) {
	if r.CreateDate == "" {
		return database.GetAllAttendanceParams{
			EmpID:      r.EmpID,
			Column2:    sql.NullTime{},
			Limit:      r.Limit,
			Offset:     r.Offset,
		}, nil
	}
	createDate, err := parseDate(r.CreateDate)
	if err != nil {
		return database.GetAllAttendanceParams{}, err
	}

	return database.GetAllAttendanceParams{
		EmpID:      r.EmpID,
		Column2:    createDate,
		CreateDate: createDate,
		Limit:      r.Limit,
		Offset:     r.Offset,
	}, nil
}


type GetAttendanceForAllReqParams struct {
	CreateDate string   `json:"create_date"`
	Limit      int32       `json:"limit"`
	Offset     int32       `json:"offset"`
}

func (r *GetAttendanceForAllReqParams) ToAllDBStruct() (database.GetAllAttendanceForAllParams, error) {
	if r.CreateDate == "" {
		return database.GetAllAttendanceForAllParams{
			Column1:    sql.NullTime{},
			Limit:      r.Limit,
			Offset:     r.Offset,
		}, nil
	}
	createDate, err := parseDate(r.CreateDate)
	if err != nil {
		return database.GetAllAttendanceForAllParams{}, err
	}

	return database.GetAllAttendanceForAllParams{
		Column1:    createDate,
		CreateDate: createDate,
		Limit:      r.Limit,
		Offset:     r.Offset,
	}, nil
}

func (r *GetAttendanceForAllReqParams) ToInsufficientDBStruct() (database.GetInsufficientAttendanceForAllParams, error) {
	if r.CreateDate == "" {
		return database.GetInsufficientAttendanceForAllParams{
			Column1:    sql.NullTime{},
			Limit:      r.Limit,
			Offset:     r.Offset,
		}, nil
	}
	createDate, err := parseDate(r.CreateDate)
	if err != nil {
		return database.GetInsufficientAttendanceForAllParams{}, err
	}

	return database.GetInsufficientAttendanceForAllParams{
		Column1:    createDate,
		CreateDate: createDate,
		Limit:      r.Limit,
		Offset:     r.Offset,
	}, nil
}

func (r *GetAttendanceForAllReqParams) ToLateDBStruct() (database.GetLateAttendanceForAllParams, error) {
	if r.CreateDate == "" {
		return database.GetLateAttendanceForAllParams{
			Column1:    sql.NullTime{},
			Limit:      r.Limit,
			Offset:     r.Offset,
		}, nil
	}
	createDate, err := parseDate(r.CreateDate)
	if err != nil {
		return database.GetLateAttendanceForAllParams{}, err
	}

	return database.GetLateAttendanceForAllParams{
		Column1:    createDate,
		CreateDate: createDate,
		Limit:      r.Limit,
		Offset:     r.Offset,
	}, nil
}

func (r *GetAttendanceForAllReqParams) ToNormalDBStruct() (database.GetNormalAttendanceForAllParams, error) {
	if r.CreateDate == "" {
		return database.GetNormalAttendanceForAllParams{
			Column1:    sql.NullTime{},
			Limit:      r.Limit,
			Offset:     r.Offset,
		}, nil
	}
	createDate, err := parseDate(r.CreateDate)
	if err != nil {
		return database.GetNormalAttendanceForAllParams{}, err
	}

	return database.GetNormalAttendanceForAllParams{
		Column1:    createDate,
		CreateDate: createDate,
		Limit:      r.Limit,
		Offset:     r.Offset,
	}, nil
}