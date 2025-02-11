package server

import (
	"github.com/Terracode-Dev/North-Star-Server/internal/database"
	hr "github.com/Terracode-Dev/North-Star-Server/internal/modules/HR"

	"github.com/labstack/echo/v4"
)

func RegisterService(e *echo.Echo, db *database.Queries) {
	hr.InitHRService(e, db).RegisterService()
}
