package hr

import (
	"log"

	db "github.com/Terracode-Dev/North-Star-Server/internal/database"
	"github.com/labstack/echo/v4"
)

type HRService struct {
	e  *echo.Echo
	db *db.Queries
}

func InitHRService(e *echo.Echo, db *db.Queries) *HRService {
	return &HRService{
		e:  e,
		db: db,
	}
}

func (S *HRService) RegisterService() {
	S.registerRoutes()
	log.Println("== HR Service Runing ==")
}
