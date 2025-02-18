package server

import (
	"database/sql"

	"github.com/Terracode-Dev/North-Star-Server/internal/database"
	hr "github.com/Terracode-Dev/North-Star-Server/internal/modules/HR"

	"github.com/labstack/echo/v4"
)

func RegisterService(e *echo.Echo, q *database.Queries, db *sql.DB) {
	hr.InitHRService(e, db, q).RegisterService()
}
