package hr

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Terracode-Dev/North-Star-Server/internal/database"
	"github.com/labstack/echo/v4"
)

type GetConfirmationReqarams struct {
	FirstName  string `json:"first_name"`
	LastName  string `json:"last_name"`
	Limit    int32       `json:"limit"`
	Offset   int32       `json:"offset"`
}

func (g *GetConfirmationReqarams) ToDbParams() (database.GetConfirmationParams, error) {
	return database.GetConfirmationParams{
		Column1: g.FirstName,
		CONCAT: g.FirstName,
		Column3: g.LastName,
		CONCAT_2: g.LastName,
		Limit: g.Limit,
		Offset: g.Offset,
	},nil
}

func(a *HRService) CreateCert(c echo.Context) error {
	emp_id, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(500, "no user id found")
	}
	err := a.q.CreateConfirmation(c.Request().Context(),int64(emp_id))
	if err != nil {
		fmt.Printf("error creating confirmation %v",err.Error())
		return c.JSON(500, "error creating confirmation")
	}
	return c.JSON(200, "confirmation created successfully")
}

func (a *HRService) GetCert(c echo.Context) error {
	var req GetConfirmationReqarams
	err := c.Bind(&req)
	if err != nil {
		fmt.Printf("error binding request %v", err.Error())
		return c.JSON(500, "error binding request")
	}
	params, err := req.ToDbParams()
	if err != nil {
		return c.JSON(500, "error converting to db struct")
	}
	cert, err := a.q.GetConfirmation(c.Request().Context(), params)
	if err != nil {
		fmt.Printf("error fetching confirmation data %v", err.Error())
		return c.JSON(500, "error fetching confirmation data")
	}
	return c.JSON(200, map[string]interface{}{
		"data": cert,
		"message": "confirmation data fetched successfully",
	})
}

func (h *HRService) DeleteCert(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("no id found", err.Error())
		return c.JSON(404, "id not found")
	}
	err = h.q.DeleteConfirmation(context.Background(), int64(id))
	if err != nil {
		return c.JSON(500, "error deleting confirmation")
	}
	return c.JSON(200, "successfully deleted")
}