package hr

import (
	"github.com/labstack/echo/v4"
)

//create tax handler
func (S *HRService) createTax(c echo.Context) error {
	var taxModel CreateTaxReqModel
	if err := c.Bind(&taxModel); err != nil {
		return c.JSON(400, err)
	}

	params, err := taxModel.ToCreateTaxParams()
	if err != nil {
		return c.JSON(400, err)
	}

	tax := S.q.CreateTax(c.Request().Context(), params)
	if tax != nil {
		return c.JSON(500, tax)
	}

	return c.JSON(200, "Tax Created")
}

//get all tax
func (S *HRService) getTax(c echo.Context) error {
	taxes, err := S.q.GetTax(c.Request().Context())
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, taxes)
}

