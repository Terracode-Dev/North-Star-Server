package hr

import (
	"log"
	"net/http"
	"time"

	db "github.com/Terracode-Dev/North-Star-Server/internal/database"
	rba "github.com/Terracode-Dev/North-Star-Server/internal/pkg/RBA"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

// @Summary Admin login
// @Description Authenticates admin and returns a JWT token in a cookie
// @Tags auth
// @Accept json
// @Produce json
// @Param user body AdminLoginReqModel true "User credentials"
// @Success 200 {string} int
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /admin/login [post]
func (S *HRService) adminLogin(c echo.Context) error {
	var loginModel AdminLoginReqModel
	if err := c.Bind(&loginModel); err != nil {
		return c.JSON(400, "request parsing error")
	}

	admin, err := S.q.AdminLogin(c.Request().Context(), loginModel.Email)
	if err != nil {
		return c.JSON(500, "internal server error")
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(loginModel.Password))
	if err != nil {
		return c.JSON(301, "invalid Password")
	}

	payload := rba.RBAauth{
		Id:     int(admin.ID),
		Role:   admin.Role,
		Email:  loginModel.Email,
		Branch: int(admin.BranchID),
	}

	t, err := rba.GenarateJWTkey(time.Hour*24, payload, []byte(S.cfg.JWTSecret))
	if err != nil {
		log.Println(err)
		return c.JSON(500, t)
	}

	cookie := new(http.Cookie)
	cookie.Name = "auth_token"
	cookie.Value = t
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(time.Hour * time.Duration(S.cfg.JwtExpHour))
	c.SetCookie(cookie)

	return c.JSON(200, "Login")
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
	var params db.SelectHrAdminParams
	if branch_id == int(S.cfg.MainBranchId) {
		params = db.SelectHrAdminParams{
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
		params = db.SelectHrAdminParams{
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
		return c.JSON(500, "internal server error")
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
