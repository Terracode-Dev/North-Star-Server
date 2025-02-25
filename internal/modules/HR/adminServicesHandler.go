package hr

import (
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// create a service
// @Summary Create Service
// @Description Create a new service
// @Tags services
// @Accept json
// @Produce json
// @Param service body CreateServicesReqModel true "Service details"
// @Success 200 {string} string "Service Created"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /service [post]
func (S *HRService) createAdminService(c echo.Context) error {
	var ser CreateServicesReqModel
	if err := c.Bind(&ser); err != nil {
		return c.JSON(400, err)
	}
	updated_by, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(400, "user not found")
	}
	updated_by64 := int64(updated_by)
	ser.UpdatedBy = &updated_by64
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
// @Summary Get Services
// @Description Get all services
// @Tags services
// @Accept json
// @Produce json
// @Success 200 {string} string "Services fetched successfully"
// @Failure 500 {string} string "internal server error"
// @Router /services [get]
func (S *HRService) getAdminServices(c echo.Context) error {
	services, err := S.q.GetServices(c.Request().Context())
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, services)
}

// get a service
// @Summary Get Service
// @Description Get a service
// @Tags services
// @Accept json
// @Produce json
// @Param category path string true "service category"
// @Success 200 {string} string "Service fetched successfully"
// @Failure 500 {string} string "internal server error"
// @Router /service/{category} [get]
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
// @Summary Update Service
// @Description Update a servicev 
func (S *HRService) updateAdminService(c echo.Context) error {
	var ser CreateServicesReqModel
	if err := c.Bind(&ser); err != nil {
		return c.JSON(400, err)
	}
	updated_by, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(400, "user not found")
	}
	updated_by64 := int64(updated_by)
	ser.UpdatedBy = &updated_by64
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
