package hr

import (
	"github.com/labstack/echo/v4"
	db "github.com/Terracode-Dev/North-Star-Server/internal/database"
)

// register admin handler
func (S *HRService) createAdmin(c echo.Context) error {
	var admin CreateHrAdminReqModel
	if err := c.Bind(&admin); err != nil {
		return c.JSON(400, err.Error())
	}
	adminParams, err := admin.convertToDbStruct()
	if err != nil {
		return c.JSON(400, err.Error())
	}
	err = S.q.CreateHrAdmin(c.Request().Context(), adminParams)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, "Admin created successfully")
}

//suspend admin handler
func (S *HRService) suspendAdmin(c echo.Context) error {
	var admin db.SuspendedHrAdminParams
	if err := c.Bind(&admin); err != nil {
		return c.JSON(400, err.Error())
	}
	err := S.q.SuspendedHrAdmin(c.Request().Context(), admin)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, "Admin suspended successfully")
}

