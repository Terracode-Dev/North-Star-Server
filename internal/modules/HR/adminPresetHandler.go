package hr

import (
	"github.com/labstack/echo/v4"
)

func (h *HRService) CreateAdminPreset(c echo.Context) error{
	var req CreateAdminPresetReqParams
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, err.Error())
	}
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
	presets, err := h.q.ListAdminPresets(c.Request().Context())
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, presets)
}