package hr

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	rba "github.com/Terracode-Dev/North-Star-Server/internal/pkg/RBA"

	"github.com/labstack/echo/v4"
)

// TestLogin godoc
// @Summary User login
// @Description Authenticates user and returns a JWT token in a cookie
// @Tags auth
// @Accept json
// @Produce json
// @Param user body TUser true "User credentials"
// @Success 200 {string} string "login"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /testlogin [post]
func (S *HRService) TestLogin(c echo.Context) error {
	u := new(TUser)
	if err := c.Bind(u); err != nil {
		log.Println("test issu")
		return err
	}
	// test only [mock data]
	payload := rba.RBAauth{
		Id:     1,
		Role:   "emp",
		Email:  "emp@mail.com",
		Branch: 2,
	}
	t, err := rba.GenarateJWTkey(time.Hour*24, payload, []byte(S.cfg.JWTSecret))
	if err != nil {
		log.Println(err)
		return c.JSON(500, t)
	}

	cookie := new(http.Cookie)
	cookie.Name = "auth_token"
	cookie.Value = t
	cookie.Expires = time.Now().Add(time.Hour * 24)
	c.SetCookie(cookie)

	return c.String(200, "login")
}

// TestAuth godoc
// @Summary User Auth
// @Description Authenticates user and returns a JWT token in a cookie
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {string} string "login"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /testauth [get]
func (S *HRService) TestAuth(c echo.Context) error {
	id := c.Get("user_id")
	branch := c.Get("branch")
	role := c.Get("role")
	return c.String(200, fmt.Sprintf("%d, %d, %s", id, branch, role))
}

// TestS3Upload godoc
// @Summary Upload a file to S3
// @Description Uploads a file to S3 and returns the filename
// @Tags upload-file
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Success 201 {string} string "File uploaded successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /testS3upload [post]
func (S *HRService) TestS3Upload(c echo.Context) error {
	data, err := c.FormFile("file")
	if err != nil {
		return c.JSON(500, "file reading issue")
	}
	fileExt := filepath.Ext(data.Filename)
	obj, err := data.Open()
	if err != nil {
		log.Println("!!! file read error [TestS3Upload] !!!\n", err)
		return c.JSON(500, "internal server error")
	}
	defer obj.Close()
	err = S.s3.UploadToS3(c.Request().Context(), "nsappvisa", data.Filename, obj)
	if err != nil {
		log.Println("!!! S3 file upload error [TestS3Upload] !!!\n", err)
		return c.JSON(500, "internal server error")
	}
	return c.String(201, fileExt)
}
