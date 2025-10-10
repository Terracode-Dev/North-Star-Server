package hr

import (
	"fmt"
	"strconv"

	"github.com/Terracode-Dev/North-Star-Server/internal/database"
	"github.com/labstack/echo/v4"
)

func (h *HRService) CreateAdminPreset(c echo.Context) error{
	var req CreateAdminPresetReqParams
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, err.Error())
	}
	fmt.Printf("Request %+v\n", req)
	params, err := req.ToCreateAdminPresetParams()
	if err != nil {
		return c.JSON(500, err.Error())
	}
	err = h.q.CreateAdminPreset(c.Request().Context(), params)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, "Preset Created Successfully")
}

func (h *HRService) GetAdminPresetBySlug(c echo.Context) error {
	slug := c.Param("slug")
	preset, err := h.q.GetAdminPresetBySlug(c.Request().Context(), slug)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, preset)
}

func (h *HRService) ListAdminPresets(c echo.Context) error {
	var req database.ListAdminPresetsParams
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, "error binding request")
	}
	presets, err := h.q.ListAdminPresets(c.Request().Context(), req)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	totalCount , err := h.q.TotalAdminPresetsCount(c.Request().Context())
	if err != nil {
		fmt.Printf("error getting total count: %v", err.Error())
		return c.JSON(500, "error getting total count")
	}
	return c.JSON(200, map[string]interface{}{
		"total_rows": totalCount,
		"presets": presets,
	})
}

func (h *HRService) DeleteAdminPreset(c echo.Context) error {
	id , err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, "invalid id")
	}
	err = h.q.DeleteAdminPresetByID(c.Request().Context(), int64(id))
	if err != nil {
		fmt.Printf("error deleting preset: %v\n", err)
		return c.JSON(500, "error deleting preset")
	}
	return c.JSON(200, "preset deleted successfully")
}

func (h *HRService) UpdateAdminPresetByID(c echo.Context) error {
	var req CreateAdminPresetReqParams
	if err := c.Bind(&req); err != nil {
		fmt.Printf("error binding request %v", err.Error())
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, "id is required")
	}
	params , err := req.ToUpdateAdminPresetParams(int64(id))
	if err != nil {
		fmt.Printf("Error converting params %v", err.Error())
		return c.JSON(500,"error converting params")
	}
	err = h.q.UpdateAdminPresetByID(c.Request().Context(), params)
	if err != nil {
		fmt.Printf("Error Updating preset: %v", err.Error())
		return c.JSON(500, "error updating preset")
	}
	return c.JSON(200, "Preset Updated Successfully")
}