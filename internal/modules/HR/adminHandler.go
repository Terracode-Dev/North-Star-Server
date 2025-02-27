package hr

import (
	"log"
	"net/http"
	"strconv"
	"time"

	db "github.com/Terracode-Dev/North-Star-Server/internal/database"
	rba "github.com/Terracode-Dev/North-Star-Server/internal/pkg/RBA"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// register admin handler

// @Summary Create Admin
// @Description Create a new admin
// @Tags admin
// @Accept json
// @Produce json
// @Param user body CreateHrAdminReqModel true "Admin details"
// @Success 200 {string} string "Admin created successfully"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /admin [post]
func (S *HRService) createAdmin(c echo.Context) error {
	var admin CreateHrAdminReqModel
	if err := c.Bind(&admin); err != nil {
		return c.JSON(400, err.Error())
	}
	updated_by, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(400, "user not found")
	}
	adminParams, err := admin.convertToDbStruct(int64(updated_by))
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
// @Router /admin/login [get]

func (S *HRService) adminLogin(c echo.Context) error {
	var loginModel AdminLoginReqModel
	if err := c.Bind(&loginModel); err != nil {
		return c.JSON(400, "request parsing error")
	}

	admin, err := S.q.AdminLogin(c.Request().Context(), loginModel.Email)
	if err != nil {
		return c.JSON(301, "invalid Email")
	}

	if admin.Status == "suspended" {
		return c.JSON(301, "Admin is suspended")
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
	cookie.HttpOnly = true
	cookie.Secure = false
	cookie.Expires = time.Now().Add(time.Hour * time.Duration(S.cfg.JwtExpHour))
	cookie.SameSite = http.SameSiteLaxMode
	c.SetCookie(cookie)

	res := AdminLoginResModel{
		Id:         int(admin.ID),
		Role:       admin.Role,
		Email:      loginModel.Email,
		Branch:     int(admin.BranchID),
		BranchName: admin.Branchname.String,
	}

	return c.JSON(200, res)
}

// get all admin handler
// @Summary Get all Admins
// @Description Get all admins
// @Tags admin
// @Accept json
// @Produce json
// @Param search query string false "search query"
// @Param pageNumber query int false "page number"
// @Param limit query int false "limit"
// @Success 200 {object} GetAdminReqModel
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /admin/all [get]
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

func (S *HRService) updateAdmin(c echo.Context) error {
	var admin CreateHrAdminReqModel
	if err := c.Bind(&admin); err != nil {
		return c.JSON(400, err.Error())
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	updated_by, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(400, "user not found")
	}
	adminParams, err := admin.convertToDbStructForUpdate(id, int64(updated_by))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	err = S.q.UpdateHrAdmin(c.Request().Context(), adminParams)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, "Admin updated successfully")
}

// suspend admin handler
// @Summary Suspend Admin
// @Description Suspend an admin
// @Tags admin
// @Accept json
// @Produce json
// @Param admin body SuspendedHrAdminParams true "Admin details"
// @Success 200 {string} string "Admin suspended successfully"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /admin/suspend [put]

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
