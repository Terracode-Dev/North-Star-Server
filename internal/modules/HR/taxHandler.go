package hr

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

// create tax handler
func (S *HRService) createTax(c echo.Context) error {
	var taxModel CreateTaxReqModel
	if err := c.Bind(&taxModel); err != nil {
		return c.JSON(400, err)
	}

	updated_by, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(301, "authentication issue")
	}
	params, err := taxModel.ToCreateTaxParams(int64(updated_by))
	if err != nil {
		return c.JSON(400, err)
	}

	tax := S.q.CreateTax(c.Request().Context(), params)
	if tax != nil {
		return c.JSON(500, tax)
	}

	return c.JSON(200, "Tax Created")
}

// get all tax
func (S *HRService) getTax(c echo.Context) error {
	taxes, err := S.q.GetTax(c.Request().Context())
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, taxes)
}

func (S *HRService) deleteTax(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}

	allowance := S.q.DeleteTax(c.Request().Context(), int64(id))
	if allowance != nil {
		return c.JSON(500, allowance)
	}

	return c.JSON(200, "Tax Deleted")
}
