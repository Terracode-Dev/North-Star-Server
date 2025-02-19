package hr

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

// create allowances
func (S *HRService) createAllowances(c echo.Context) error {
	var allow CreateAllowancesReqModel
	if err := c.Bind(&allow); err != nil {
		return c.JSON(400, err)
	}

	params, err := allow.ToCreateAllowancesParams()
	if err != nil {
		return c.JSON(400, err)
	}

	allowance := S.q.CreateAllowances(c.Request().Context(), params)
	if allowance != nil {
		return c.JSON(500, allowance)
	}

	return c.JSON(200, "Allowance Created")
}

// get all allowance
func (S *HRService) getAllowances(c echo.Context) error {
	allowances, err := S.q.GetAllowances(c.Request().Context())
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, allowances)
}

// get a allowance
func (S *HRService) getOneAllowance(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}

	allowance, err := S.q.GetAllowance(c.Request().Context(), int64(id))
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, allowance)
}

// update a allowance
func (S *HRService) updateAllowance(c echo.Context) error {
	var allow CreateAllowancesReqModel
	if err := c.Bind(&allow); err != nil {
		return c.JSON(400, err)
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}

	params := allow.ToUpdateAllowancesParams(int64(id))
	allowance := S.q.UpdateAllowance(c.Request().Context(), params)
	if allowance != nil {
		return c.JSON(500, allowance)
	}

	return c.JSON(200, "Allowance Updated")
}

// delete a allowance
func (S *HRService) deleteAllowance(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}

	allowance := S.q.DeleteAllowance(c.Request().Context(), int64(id))
	if allowance != nil {
		return c.JSON(500, allowance)
	}

	return c.JSON(200, "Allowance Deleted")
}
