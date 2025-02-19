package hr

import (
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// create a service
func (S *HRService) createAdminService(c echo.Context) error {
	var ser CreateServicesReqModel
	if err := c.Bind(&ser); err != nil {
		return c.JSON(400, err)
	}

	params, err := ser.ToCreateServicesParams()
	if err != nil {
		return c.JSON(400, err)
	}

	service := S.q.CreateServices(c.Request().Context(), params)
	if service != nil {
		return c.JSON(500, service)
	}

	return c.JSON(200, "Service Created")
}

// get all services
func (S *HRService) getAdminServices(c echo.Context) error {
	services, err := S.q.GetServices(c.Request().Context())
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, services)
}

// get a service
func (S *HRService) getOneAdminService(c echo.Context) error {
	cat := c.Param("category")
	cat = strings.ReplaceAll(cat, "_", " ")

	service, err := S.q.GetService(c.Request().Context(), cat)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, service)
}

// update a service
func (S *HRService) updateAdminService(c echo.Context) error {
	var ser CreateServicesReqModel
	if err := c.Bind(&ser); err != nil {
		return c.JSON(400, err)
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}

	params := ser.ToUpdateServicesParams(int64(id))
	service := S.q.UpdateService(c.Request().Context(), params)
	if service != nil {
		return c.JSON(500, service)
	}

	return c.JSON(200, "Service Updated")

}

// delete a service
func (S *HRService) deleteAdminService(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}

	service := S.q.DeleteService(c.Request().Context(), int64(id))
	if service != nil {
		return c.JSON(500, service)
	}

	return c.JSON(200, "Service Deleted")
}
