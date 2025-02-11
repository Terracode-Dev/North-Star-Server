package hr

import "github.com/labstack/echo/v4"

func (S *HRService) createAdmin(c echo.Context) error {
	return c.String(200, "test")
}
