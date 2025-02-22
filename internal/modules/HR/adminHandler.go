package hr

import (
	"github.com/Terracode-Dev/North-Star-Server/internal/database"
	db "github.com/Terracode-Dev/North-Star-Server/internal/database"
	"github.com/labstack/echo/v4"
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

func (S *HRService) getAllAdmin(c echo.Context) error {
	var req GetAdminReqModel
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, err.Error())
	}

	branch_id, ok := c.Get("branch").(int)
	if !ok {
		return c.JSON(400, "Invalid branch ID")
	}
	var params database.SelectHrAdminParams
	if branch_id == int(S.cfg.MainBranchId) {
		params = database.SelectHrAdminParams{
			CONCAT:   req.Search,
			CONCAT_2: req.Search,
			CONCAT_3: req.Search,
			CONCAT_4: req.Search,
			Column5:  "",
			BranchID: 0,
			Limit:    req.Limit,
			Offset:   (req.PageNumber - 1) * req.Limit,
		}
	} else {
		params = database.SelectHrAdminParams{
			CONCAT:   req.Search,
			CONCAT_2: req.Search,
			CONCAT_3: req.Search,
			CONCAT_4: req.Search,
			Column5:  "1",
			BranchID: int64(branch_id),
			Limit:    req.Limit,
			Offset:   (req.PageNumber - 1) * req.Limit,
		}
	}

	admins, err := S.q.SelectHrAdmin(c.Request().Context(), params)
	if err != nil {
		return c.JSON(500, "Error getting employee")
	}
	return c.JSON(200, admins)
}

// suspend admin handler
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
