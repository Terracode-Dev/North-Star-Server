package hr

import (
	"log"

	"github.com/labstack/echo/v4"
)

type VerifyDataModel struct {
	Id         int    `json:"id"`
	Role       string `json:"role"`
	Email      string `json:"email"`
	Branch     int    `json:"branch"`
	BranchName string `json:"branchName"`
}

func (S *HRService) verifyAuth(c echo.Context) error {
	email, ok := c.Get("email").(string)
	if !ok {
		log.Println("LOGAUTH", ok)
		return c.JSON(401, map[string]interface{}{"isAuthenticated": false, "message": "route not accessible"})
	}

	admin, err := S.q.AdminLogin(c.Request().Context(), email)
	if err != nil {
		log.Println("LOGAUTH", err)
		return c.JSON(401, map[string]interface{}{"isAuthenticated": false, "message": "invalid Email"})
	}

	if !admin.Branchname.Valid {
		log.Println("LOGAUTH", "saaa")
		return c.JSON(401, map[string]interface{}{"isAuthenticated": false, "message": "route not accessible"})
	}

	res := VerifyDataModel{
		Id:         int(admin.ID),
		Role:       admin.Role,
		Email:      email,
		Branch:     int(admin.BranchID),
		BranchName: admin.Branchname.String,
	}

	log.Println(res)

	// Include `isAuthenticated` in the response
	return c.JSON(200, map[string]interface{}{
		"isAuthenticated": true,
		"data":            res,
	})
}

func (S *HRService) empVerifyAuth(c echo.Context) error {
	email, ok := c.Get("email").(string)
	if !ok {
		log.Println("LOGAUTH", ok)
		return c.JSON(401, map[string]interface{}{"isAuthenticated": false, "message": "route not accessible"})
	}

	user_id, ok := c.Get("user_id").(int)
	if !ok {
		log.Println("LOGAUTH", ok)
		return c.JSON(401, map[string]interface{}{"isAuthenticated": false, "message": "route not accessible"})
	}

	res := VerifyDataModel{
		Id:    user_id,
		Email: email,
		Role:  "emp",
	}

	return c.JSON(200, res)
}
