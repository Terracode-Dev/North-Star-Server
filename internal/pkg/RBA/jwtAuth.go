package rba

import (
	"net/http"

	"github.com/Terracode-Dev/North-Star-Server/internal/config"
	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
)

type RBAauth struct {
	Id    int    `json:"id"`
	Role  string `json:"role"`
	Email string `json:"email"`
}

type JWTPayload struct {
	Data RBAauth `json:"data"`
	jwt.RegisteredClaims
}

func AuthMiddelware(role []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sec := config.LoadConfig().JWTSecret
			t, err := c.Cookie("auth_token")
			if err != nil {
				return err
			}
			data, err := ValidateJWTkey(t.Value, []byte(sec))
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "unauthorized route access")
			}
			tokenRole := data.Data.Role
			if contains(role, tokenRole) {
				c.Set("user_id", data.Data.Id)
				return next(c)
			}
			return c.JSON(http.StatusUnauthorized, "unauthorized route access")
		}
	}
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// func deleteCookie(c echo.Context) error {
// 	// Create a new cookie with the same name, empty value, and expiration in the past
// 	cookie := new(http.Cookie)
// 	cookie.Name = "session_token"
// 	cookie.Value = ""
// 	cookie.Expires = time.Unix(0, 0) // Expire the cookie immediately
// 	cookie.MaxAge = -1               // Force the browser to delete it
// 	cookie.Path = "/"                // Ensure it's deleted from the entire site
//
// 	// Set the cookie to the response
// 	c.SetCookie(cookie)
//
// 	return c.String(http.StatusOK, "Cookie deleted")
// }
