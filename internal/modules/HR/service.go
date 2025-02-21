package hr

import (
	"database/sql"
	"log"

	"github.com/Terracode-Dev/North-Star-Server/internal/config"
	"github.com/Terracode-Dev/North-Star-Server/internal/database"
	aws "github.com/Terracode-Dev/North-Star-Server/internal/pkg/aws"

	"github.com/labstack/echo/v4"
)

type HRService struct {
	e   *echo.Echo
	cfg *config.Config
	db  *sql.DB
	q   *database.Queries
	s3  *aws.S3Client
}

func InitHRService(e *echo.Echo, cfg *config.Config, db *sql.DB, q *database.Queries, s3client *aws.S3Client) *HRService {
	return &HRService{
		e:   e,
		cfg: cfg,
		q:   q,
		db:  db,
		s3:  s3client,
	}
}

func (S *HRService) RegisterService() {
	S.registerRoutes()
	log.Println("== HR Service Runing ==")
}
