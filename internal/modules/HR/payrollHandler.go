package hr

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

// create payroll handler
// @Summary Create Payroll
// @Description Create a new payroll
// @Tags payroll
// @Accept json
// @Produce json
// @Param payroll body PayrollAllowances true "Payroll details"
// @Success 200 {string} string "Payroll Created"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /payroll [post]
func (S *HRService) createPayroll(c echo.Context) error {
	var pay PayrollAllowances
	if err := c.Bind(&pay); err != nil {
		return c.JSON(400, err.Error())
	}
	tx, err := S.db.Begin()
	if err != nil {
		return c.JSON(500, "Error starting transaction")
	}
	defer tx.Rollback()
	qtx := S.q.WithTx(tx)

	updated_by, ok:= c.Get("user_id").(int)
	if !ok {
		return c.JSON(400, "user not found")
	}

	paytollParams, err := pay.Payroll.ToCreatePayrollParams(int64(updated_by))
	if err != nil {
		return c.JSON(400, err.Error())
	}

	payroll, err:= qtx.CreatePayroll(c.Request().Context(), paytollParams)
	if err != nil {
		return c.JSON(500, "Error creating payroll")
	}

	payroll_id, err := payroll.LastInsertId()
	if err != nil {
		return c.JSON(500, "Error getting payroll id")
	}

	for _, allowance := range pay.Allowances {
		allowance.PayrollID = payroll_id
		allowancesParams, err := allowance.ToCreatePayrollAllowancesParams(int64(updated_by))
		if err != nil {
			return c.JSON(400, err.Error())
		}
		
		err = qtx.CreatePayrollAllowances(c.Request().Context(), allowancesParams)
		if err != nil {
			return c.JSON(500, "Error creating payroll allowances")
		}
	}

	err = tx.Commit()
	if err != nil {
		return c.JSON(500, "Error committing transaction")
	}

	return c.JSON(200, "Payroll Created")

}

// get all payroll
// @Summary Get Payrolls
// @Description Get all payrolls
// @Tags payroll
// @Accept json
// @Produce json
// @Param search query string false "search query"
// @Param pageNumber query int false "page number"
// @Param limit query int false "limit"
// @Success 200 {object} GetPayrollsReqModel
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /payroll [get]
func (S *HRService) getPayroll(c echo.Context) error {
	var payModel GetPayrollsReqModel
	if err := c.Bind(&payModel); err != nil {
		return c.JSON(400, err.Error())
	}
	payParams, err := payModel.ToGetPayrollsParams()
	if err != nil {
		return c.JSON(400, err.Error())
	}
	payroll, err := S.q.GetPayrolls(c.Request().Context(),payParams)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, payroll)
}

// get one payroll
// @Summary Get Payroll
// @Description Get one payroll
// @Tags payroll
// @Accept json
// @Produce json
// @Param id path int true "payroll id"
// @Success 200 {object} PayrollAllowances
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /payroll/{id} [get]
func (S *HRService) getOnePayroll(c echo.Context) error {
	payId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	payroll, err := S.q.GetOnePayroll(c.Request().Context(), payId)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, payroll)
}

// update payroll
// @Summary Update Payroll
// @Description Update payroll
// @Tags payroll
// @Accept json
// @Produce json
// @Param id path int true "payroll id"
// @Param payroll body CreatePayrollReqModel true "Payroll details"
// @Success 200 {string} string "Payroll Updated"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /payroll/{id} [put]
func (S *HRService) updatePayroll(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	updated_by := c.Get("user").(int)
	var pay CreatePayrollReqModel
	if err := c.Bind(&pay); err != nil {
		return c.JSON(400, err.Error())
	}
	payrollParams, err := pay.ToUpdatePayrollParams(int64(id), int64(updated_by))
	if err != nil {	
		return c.JSON(400, err.Error())
	}
	err = S.q.UpdatePayroll(c.Request().Context(), payrollParams)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, "Payroll Updated")
}

