package hr

import (
	"database/sql"
	"log"

	"github.com/Terracode-Dev/North-Star-Server/internal/database"
	db "github.com/Terracode-Dev/North-Star-Server/internal/database"
	"github.com/labstack/echo/v4"
)

type HRService struct {
	e  *echo.Echo
	db *sql.DB
	q *db.Queries
}

func InitHRService(e *echo.Echo, db *sql.DB, q *database.Queries) *HRService {
	return &HRService{
		e:  e,
		q: q,
		db: db,
	}
}

func (S *HRService) RegisterService() {
	S.registerRoutes()
	log.Println("== HR Service Runing ==")
}
