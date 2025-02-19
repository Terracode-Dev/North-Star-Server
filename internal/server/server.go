package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Terracode-Dev/North-Star-Server/internal/config"
	"github.com/Terracode-Dev/North-Star-Server/internal/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func InitServer() {
	cfg := config.LoadConfig()
	// Setup
	e := echo.New()
	queries, db := database.CreateNewDB(cfg.DBString)
	e.Logger.SetLevel(log.INFO)

	// middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// NOTE: Only For Developmet

	// CORS default
	// Allows requests from any origin wth GET, HEAD, PUT, POST or DELETE method.
	e.Use(middleware.CORS())

	// service Registation
	RegisterService(e, queries, db)

	e.GET("/", func(c echo.Context) error {
		time.Sleep(5 * time.Second)
		return c.JSON(http.StatusOK, "OK")
	})

	s := &http.Server{
		Addr:         cfg.Port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	e.Server = s

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.StartServer(s); err != nil && err != http.ErrServerClosed {
			log.Print(err)
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	<-ctx.Done()
	e.Logger.Info("Received interrupt signal, shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	e.Logger.Info("Server gracefully stopped")
}
