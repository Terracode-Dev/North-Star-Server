package hr

import (
	"database/sql"
	"strconv"
	"net/http"
	"github.com/labstack/echo/v4"
    "github.com/Terracode-Dev/North-Star-Server/internal/database"
)

func (S *HRService) SalaryTransfer(c echo.Context) error {
    var req GetAccountDetailsReqParams
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    dateParams, err := req.ConvertToDbParams()
    if err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    data, err := S.q.GetAccountDetails(c.Request().Context(), dateParams)
    if err != nil {
        if err == sql.ErrNoRows {
            return c.JSON(http.StatusNotFound, "No account details found for the given date")
        }
        return c.JSON(http.StatusInternalServerError, "Error fetching account details")
    }
    var total float64 = 0.0
    for _, account := range data {
        paidAmount, err := strconv.ParseFloat(account.AmountPaid.String(), 64)
        if err != nil {
            paidAmount = 0.0
        }
        total += paidAmount
    }
    response := SalaryTransferRes{
        AccountData: data, 
        Total:       total,
    }
    return c.JSON(http.StatusOK, response)
}

func (S *HRService) GetExpiredVisaOrReports(c echo.Context) error {
    var req database.GetExpiredVisaOrReportsParams
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
	data, err := S.q.GetExpiredVisaOrReports(c.Request().Context(), req)
    if err != nil {
        if err == sql.ErrNoRows {
            return c.JSON(http.StatusNotFound, "No expired visa or reports found")
        }
        return c.JSON(http.StatusInternalServerError, "Error fetching expired visa or reports")
    }
    return c.JSON(http.StatusOK, data)
}

func (S *HRService) GetSoonExpiringPassportsAndReports(c echo.Context) error {
    var req database.GetVisaOrPassportExpiringSoonParams
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    data, err := S.q.GetVisaOrPassportExpiringSoon(c.Request().Context(), req)
    if err != nil {
        if err == sql.ErrNoRows {
            return c.JSON(http.StatusNotFound, "No soon expiring passports or reports found")
        }
        return c.JSON(http.StatusInternalServerError, "Error fetching soon expiring passports or reports")
    }
    return c.JSON(http.StatusOK, data)
}

func (S *HRService) GetStaffPayroll(c echo.Context) error {
    var req GetStaffPayrollReqParams
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    dateParams, err := req.ConvertToDbParams()
    if err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    data, err := S.q.GetStaffPayroll(c.Request().Context(), dateParams)
    if err != nil {
        if err == sql.ErrNoRows {
            return c.JSON(http.StatusNotFound, "No staff payroll found for the given date")
        }
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, data)
}

func (S *HRService) GetemployeeInsurance(c echo.Context) error {
    branchID, err := strconv.ParseInt(c.Param("branch_id"), 10, 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid branch ID")
    }
    data , err := S.q.GetempployeeInsurance(c.Request().Context() , branchID)
    if err != nil {
        if err == sql.ErrNoRows {
            return c.JSON(http.StatusNotFound, "No employee insurance found")
        }
        return c.JSON(http.StatusInternalServerError, "Error fetching employee insurance")
    }
    return c.JSON(http.StatusOK, data)
}