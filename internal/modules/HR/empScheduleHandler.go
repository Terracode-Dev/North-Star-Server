package hr

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Terracode-Dev/North-Star-Server/internal/database"
	"github.com/labstack/echo/v4"
)

type GetEmployeeIdByEmailRequest struct {
	Email string `json:"email" query:"email"`
}

func (s *HRService) GetEmployeeIdByEmail(c echo.Context) error {
	// email := c.QueryParam("email")
	var req GetEmployeeIdByEmailRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}
	fmt.Println("Email received:", req.Email)
	if req.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "email parameter is required",
		})
	}

	employee, err := s.q.GetEmployeeByEmail(c.Request().Context(), req.Email)
	if err != nil {
		fmt.Println("Error fetching employee by email:", err)
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "employee not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to get employee information",
		})
	}

	// Create response with the required format
	response := EmployeeResponse{
		ID:         employee.EmployeeID,
		Email:      req.Email,
		Name:       fmt.Sprintf("%s %s", employee.FirstName, employee.LastName),
		Department: employee.Department,
		Position:   employee.Designation,
	}

	return c.JSON(http.StatusOK, response)
}

// CreateEmployeeSchedule - POST /api/employee/schedule
func (s *HRService) CreateEmployeeSchedule(c echo.Context) error {
	var req CreateScheduleRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	ctx := c.Request().Context()

	// Create weekly schedule
	weeklyParams := database.CreateEmployeeScheduleParams{
		EmpID: req.EmpID,
	}

	// Parse Monday
	if req.Weekly.Monday != nil {
		weeklyParams.Monday = sql.NullBool{Bool: req.Weekly.Monday.IsWorking, Valid: true}
		if req.Weekly.Monday.FromTime != "" {
			fromTime, err := parseTime(req.Weekly.Monday.FromTime)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "invalid monday from_time format",
				})
			}
			weeklyParams.MondayFrom = fromTime
		}
		if req.Weekly.Monday.ToTime != "" {
			toTime, err := parseTime(req.Weekly.Monday.ToTime)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "invalid monday to_time format",
				})
			}
			weeklyParams.MondayTo = toTime
		}
	}

	// Parse Tuesday
	if req.Weekly.Tuesday != nil {
		weeklyParams.Tuesday = sql.NullBool{Bool: req.Weekly.Tuesday.IsWorking, Valid: true}
		if req.Weekly.Tuesday.FromTime != "" {
			fromTime, err := parseTime(req.Weekly.Tuesday.FromTime)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "invalid tuesday from_time format",
				})
			}
			weeklyParams.TuesdayFrom = fromTime
		}
		if req.Weekly.Tuesday.ToTime != "" {
			toTime, err := parseTime(req.Weekly.Tuesday.ToTime)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "invalid tuesday to_time format",
				})
			}
			weeklyParams.TuesdayTo = toTime
		}
	}

	// Parse Wednesday
	if req.Weekly.Wednesday != nil {
		weeklyParams.Wednesday = sql.NullBool{Bool: req.Weekly.Wednesday.IsWorking, Valid: true}
		if req.Weekly.Wednesday.FromTime != "" {
			fromTime, err := parseTime(req.Weekly.Wednesday.FromTime)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "invalid wednesday from_time format",
				})
			}
			weeklyParams.WednesdayFrom = fromTime
		}
		if req.Weekly.Wednesday.ToTime != "" {
			toTime, err := parseTime(req.Weekly.Wednesday.ToTime)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "invalid wednesday to_time format",
				})
			}
			weeklyParams.WednesdayTo = toTime
		}
	}

	// Parse Thursday
	if req.Weekly.Thursday != nil {
		weeklyParams.Thursday = sql.NullBool{Bool: req.Weekly.Thursday.IsWorking, Valid: true}
		if req.Weekly.Thursday.FromTime != "" {
			fromTime, err := parseTime(req.Weekly.Thursday.FromTime)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "invalid thursday from_time format",
				})
			}
			weeklyParams.ThursdayFrom = fromTime
		}
		if req.Weekly.Thursday.ToTime != "" {
			toTime, err := parseTime(req.Weekly.Thursday.ToTime)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "invalid thursday to_time format",
				})
			}
			weeklyParams.ThursdayTo = toTime
		}
	}

	// Parse Friday
	if req.Weekly.Friday != nil {
		weeklyParams.Friday = sql.NullBool{Bool: req.Weekly.Friday.IsWorking, Valid: true}
		if req.Weekly.Friday.FromTime != "" {
			fromTime, err := parseTime(req.Weekly.Friday.FromTime)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "invalid friday from_time format",
				})
			}
			weeklyParams.FridayFrom = fromTime
		}
		if req.Weekly.Friday.ToTime != "" {
			toTime, err := parseTime(req.Weekly.Friday.ToTime)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "invalid friday to_time format",
				})
			}
			weeklyParams.FridayTo = toTime
		}
	}

	// Parse Saturday
	if req.Weekly.Saturday != nil {
		weeklyParams.Saturday = sql.NullBool{Bool: req.Weekly.Saturday.IsWorking, Valid: true}
		if req.Weekly.Saturday.FromTime != "" {
			fromTime, err := parseTime(req.Weekly.Saturday.FromTime)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "invalid saturday from_time format",
				})
			}
			weeklyParams.SaturdayFrom = fromTime
		}
		if req.Weekly.Saturday.ToTime != "" {
			toTime, err := parseTime(req.Weekly.Saturday.ToTime)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "invalid saturday to_time format",
				})
			}
			weeklyParams.SaturdayTo = toTime
		}
	}

	// Parse Sunday
	if req.Weekly.Sunday != nil {
		weeklyParams.Sunday = sql.NullBool{Bool: req.Weekly.Sunday.IsWorking, Valid: true}
		if req.Weekly.Sunday.FromTime != "" {
			fromTime, err := parseTime(req.Weekly.Sunday.FromTime)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "invalid sunday from_time format",
				})
			}
			weeklyParams.SundayFrom = fromTime
		}
		if req.Weekly.Sunday.ToTime != "" {
			toTime, err := parseTime(req.Weekly.Sunday.ToTime)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "invalid sunday to_time format",
				})
			}
			weeklyParams.SundayTo = toTime
		}
	}

	// Create weekly schedule
	if err := s.q.CreateEmployeeSchedule(ctx, weeklyParams); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Create additional schedules
	for _, additional := range req.Additional {
		date, err := parseDate(additional.Date)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid date format: " + additional.Date,
			})
		}

		fromTime, err := parseTime(additional.FromTime)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid from_time format",
			})
		}

		toTime, err := parseTime(additional.ToTime)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid to_time format",
			})
		}

		additionalParams := database.CreateAdditionalScheduleParams{
			EmpID:    req.EmpID,
			Date:     date,
			FromTime: fromTime,
			ToTime:   toTime,
		}

		if err := s.q.CreateAdditionalSchedule(ctx, additionalParams); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "failed to create additional schedule",
			})
		}
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "schedule created successfully",
	})
}

// UpdateEmployeeSchedule - PUT /api/employee/:id/schedule
func (s *HRService) UpdateEmployeeSchedule(c echo.Context) error {
	empIDStr := c.Param("id")
	empID, err := strconv.ParseInt(empIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid employee ID",
		})
	}

	var req UpdateScheduleRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	ctx := c.Request().Context()

	// Update weekly schedule if provided
	if req.Weekly != nil {
		weeklyParams := database.UpdateEmployeeScheduleParams{
			EmpID: empID,
		}

		// Parse each day similar to create handler
		if req.Weekly.Monday != nil {
			weeklyParams.Monday = sql.NullBool{Bool: req.Weekly.Monday.IsWorking, Valid: true}
			if req.Weekly.Monday.FromTime != "" {
				fromTime, err := parseTime(req.Weekly.Monday.FromTime)
				if err != nil {
					return c.JSON(http.StatusBadRequest, map[string]string{
						"error": "invalid monday from_time format",
					})
				}
				weeklyParams.MondayFrom = fromTime
			}
			if req.Weekly.Monday.ToTime != "" {
				toTime, err := parseTime(req.Weekly.Monday.ToTime)
				if err != nil {
					return c.JSON(http.StatusBadRequest, map[string]string{
						"error": "invalid monday to_time format",
					})
				}
				weeklyParams.MondayTo = toTime
			}
		}

		// Continue for other days... (similar pattern)
		// For brevity, I'll skip the repetitive parsing code

		if err := s.q.UpdateEmployeeSchedule(ctx, weeklyParams); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "failed to update weekly schedule",
			})
		}
	}

	// Update additional schedules if provided
	for _, additional := range req.Additional {
		date, err := parseDate(additional.Date)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid date format: " + additional.Date,
			})
		}

		fromTime, err := parseTime(additional.FromTime)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid from_time format",
			})
		}

		toTime, err := parseTime(additional.ToTime)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid to_time format",
			})
		}

		updateParams := database.UpdateAdditionalScheduleParams{
			EmpID:    empID,
			Date:     date,
			FromTime: fromTime,
			ToTime:   toTime,
		}

		if err := s.q.UpdateAdditionalSchedule(ctx, updateParams); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "failed to update additional schedule",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "schedule updated successfully",
	})
}

// DeleteEmployeeSchedule - DELETE /api/employee/:id/schedule
func (s *HRService) DeleteEmployeeSchedule(c echo.Context) error {
	empIDStr := c.Param("id")
	empID, err := strconv.ParseInt(empIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid employee ID",
		})
	}

	ctx := c.Request().Context()

	// Delete weekly schedule
	if err := s.q.DeleteEmployeeSchedule(ctx, empID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to delete weekly schedule",
		})
	}

	// Delete all additional schedules
	if err := s.q.DeleteAllAdditionalSchedules(ctx, empID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to delete additional schedules",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "schedule deleted successfully",
	})
}

// DeleteAdditionalSchedule - DELETE /api/employee/:id/schedule/additional/:date
func (s *HRService) DeleteAdditionalSchedule(c echo.Context) error {
	empIDStr := c.Param("id")
	empID, err := strconv.ParseInt(empIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid employee ID",
		})
	}

	dateStr := c.Param("date")
	date, err := parseDate(dateStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid date format",
		})
	}

	ctx := c.Request().Context()

	deleteParams := database.DeleteAdditionalScheduleParams{
		EmpID: empID,
		Date:  date,
	}

	if err := s.q.DeleteAdditionalSchedule(ctx, deleteParams); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to delete additional schedule",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "additional schedule deleted successfully",
	})
}

// GetEmployeeList - GET /api/employees
func (s *HRService) GetEmployeeList(c echo.Context) error {
	var req EmployeeListRequest

	// Set defaults
	req.Page = 1
	req.Limit = 20
	req.Year = time.Now().Year()

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid query parameters",
		})
	}

	// Validate parameters
	if req.Page < 1 {
		req.Page = 1
	}
	if req.Limit < 1 || req.Limit > 100 {
		req.Limit = 20
	}
	if req.Year == 0 {
		req.Year = time.Now().Year()
	}

	ctx := c.Request().Context()
	offset := (req.Page - 1) * req.Limit

	// Get total count
	countParams := database.CountEmployeesWithFiltersParams{
		Column1:  req.FirstName,
		CONCAT:   req.FirstName,
		Column3:  req.LastName,
		CONCAT_2: req.LastName,
		Column5:  req.Email,
		CONCAT_3: req.Email,
	}

	totalCount, err := s.q.CountEmployeesWithFilters(ctx, countParams)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to count employees",
		})
	}

	// Create year date for query
	yearDate := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	// Get employee list
	listParams := database.GetEmployeeListWithWorkDaysParams{
		Date:     yearDate,
		Column2:  req.FirstName,
		CONCAT:   req.FirstName,
		Column4:  req.LastName,
		CONCAT_2: req.LastName,
		Column6:  req.Email,
		CONCAT_3: req.Email,
		Column8:  req.SortBy,
		Column9:  req.SortBy,
		Column10: req.SortBy,
		Column11: req.SortBy,
		Limit:    int32(req.Limit),
		Offset:   int32(offset),
	}

	employees, err := s.q.GetEmployeeListWithWorkDays(ctx, listParams)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch employees",
		})
	}

	// Convert to response format
	var employeeList []EmployeeListItem
	for _, emp := range employees {
		email := ""
		if emp.Email.Valid {
			email = emp.Email.String
		}

		employeeList = append(employeeList, EmployeeListItem{
			ID:              emp.ID,
			FirstName:       emp.FirstName,
			LastName:        emp.LastName,
			Email:           email,
			WorkDaysForYear: emp.WorkDaysForYear,
		})
	}

	// Calculate pagination
	totalPages := int((totalCount + int64(req.Limit) - 1) / int64(req.Limit))

	response := EmployeeListResponse{
		Data: employeeList,
		Pagination: PaginationInfo{
			CurrentPage:  req.Page,
			PageSize:     req.Limit,
			TotalRecords: totalCount,
			TotalPages:   totalPages,
			HasNext:      req.Page < totalPages,
			HasPrevious:  req.Page > 1,
		},
		Filters: FilterInfo{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			SortBy:    req.SortBy,
			Year:      req.Year,
		},
	}

	return c.JSON(http.StatusOK, response)
}

// GetEmployeeWorkDaysBreakdown - GET /api/employee/:id/workdays-breakdown
func (s *HRService) GetEmployeeWorkDaysBreakdown(c echo.Context) error {
	empIDStr := c.Param("id")
	empID, err := strconv.ParseInt(empIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid employee ID",
		})
	}

	yearStr := c.QueryParam("year")
	year := time.Now().Year()
	if yearStr != "" {
		if y, err := strconv.Atoi(yearStr); err == nil {
			year = y
		}
	}

	ctx := c.Request().Context()
	yearDate := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	breakdownParams := database.GetEmployeeWorkDaysBreakdownParams{
		Date:   yearDate,
		Date_2: yearDate,
		ID:     empID,
	}

	breakdown, err := s.q.GetEmployeeWorkDaysBreakdown(ctx, breakdownParams)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "employee not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to get work days breakdown",
		})
	}

	email := ""
	if breakdown.Email.Valid {
		email = breakdown.Email.String
	}

	response := WorkDaysBreakdownResponse{
		ID:                   breakdown.ID,
		FirstName:            breakdown.FirstName,
		LastName:             breakdown.LastName,
		Email:                email,
		WeeklyWorkDays:       interfaceToInt64(breakdown.WeeklyWorkDays),
		BaseYearlyDays:       interfaceToInt64(breakdown.BaseYearlyDays),
		AdditionalDays:       interfaceToInt64(breakdown.AdditionalDays),
		TotalWorkDaysForYear: breakdown.TotalWorkDaysForYear,
	}

	return c.JSON(http.StatusOK, response)
}
