package hr

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

// create payroll handler
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

	paytollParams, err := pay.Payroll.ToCreatePayrollParams()
	if err != nil {
		return c.JSON(400, err.Error())
	}

	payroll, err:= qtx.CreatePayroll(c.Request().Context(), paytollParams)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	payroll_id, err := payroll.LastInsertId()
	if err != nil {
		return c.JSON(500, err.Error())
	}

	for _, allowance := range pay.Allowances {
		allowance.PayrollID = payroll_id
		allowancesParams, err := allowance.ToCreatePayrollAllowancesParams()
		if err != nil {
			return c.JSON(400, err.Error())
		}
		
		err = qtx.CreatePayrollAllowances(c.Request().Context(), allowancesParams)
		if err != nil {
			return c.JSON(500, err.Error())
		}
	}

	err = tx.Commit()
	if err != nil {
		return c.JSON(500, "Error committing transaction")
	}

	return c.JSON(200, "Payroll Created")

}

// get all payroll
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
func (S *HRService) updatePayroll(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err.Error())
	}

	var pay CreatePayrollReqModel
	if err := c.Bind(&pay); err != nil {
		return c.JSON(400, err.Error())
	}
	payrollParams, err := pay.ToUpdatePayrollParams(int64(id))
	if err != nil {	
		return c.JSON(400, err.Error())
	}
	err = S.q.UpdatePayroll(c.Request().Context(), payrollParams)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, "Payroll Updated")
}

