package server

import (
	"database/sql"

	"github.com/Terracode-Dev/North-Star-Server/internal/config"
	"github.com/Terracode-Dev/North-Star-Server/internal/database"
	hr "github.com/Terracode-Dev/North-Star-Server/internal/modules/HR"
	aws "github.com/Terracode-Dev/North-Star-Server/internal/pkg/aws"

	"github.com/labstack/echo/v4"
)

func RegisterService(e *echo.Echo, cfg *config.Config, q *database.Queries, db *sql.DB, s3 *aws.S3Client) {
	hr.InitHRService(e, cfg, db, q, s3).RegisterService()
}
