package hr

import (
	"strconv"
	"github.com/labstack/echo/v4"
)

func (h *HRService) CreateEmpLink(c echo.Context) error{
	var req CreateEmpLinkReqParams
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, err.Error())
	}
	params, err := req.ToCreateEmpLinkParams()
	if err != nil {
		return c.JSON(500, err.Error())
	}
	err = h.q.CreateEmpLink(c.Request().Context(), params)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, "Employee Link Created Successfully")
}

func (h *HRService) GetEmpLinkByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	data, err := h.q.GetEmpLinkByID(c.Request().Context(), int64(id))
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, data)
}

func (h *HRService) ListEmpLinks(c echo.Context) error {
	links, err := h.q.ListEmpLinks(c.Request().Context())
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, links)
}

func (h *HRService) ApproveEmpLink(c echo.Context) error {
	var req UpdateEmpLinkApprovalReqParams
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, err.Error())
	}
	updated_by, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(301, "Authentication issue")
	}
	params, err := req.ToUpdateEmpLinkApprovalParams(int64(updated_by))
	if err != nil {
		return c.JSON(500, err.Error())
	}
	err = h.q.UpdateEmpLinkApproval(c.Request().Context(), params)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, "Employee Link Approval Updated Successfully")
}