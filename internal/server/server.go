package server

import (
	"log"
	"net/http"

	"github.com/LanternNassi/IMSController/internal/models"

	"github.com/labstack/echo"

	interfaces "github.com/LanternNassi/IMSController/internal/Interfaces"
)

type EchoServer struct {
	echo *echo.Echo
	DB   interfaces.DataBaseClient
}

func NewEchoServer(db interfaces.DataBaseClient) interfaces.Server {
	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}

	server.registerRoutes()
	return server

}

func (s *EchoServer) Start() error {
	if err := s.echo.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server shutdown occurred: %s", err)
		return err
	}

	return nil
}

func (s *EchoServer) Readiness(ctx echo.Context) error {
	ready := s.DB.Ready()
	if ready {
		return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
	}

	return ctx.JSON(http.StatusInternalServerError, models.Health{Status: "Failure"})
}

func (s *EchoServer) Liveness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})

}

func (s *EchoServer) registerRoutes() {
	s.echo.GET("/readiness", s.Readiness)
	s.echo.GET("/liveness", s.Liveness)

	cg := s.echo.Group("/clients")
	cg.GET("", s.GetClients)
	cg.POST("", s.AddClient)
	cg.PUT("/:id", s.UpdateClient)
	cg.GET("/:id", s.GetClientById)

	bg := s.echo.Group("/backup")
	bg.GET("", s.Getbackups)
	bg.POST("", s.AddBackup)
	bg.GET("/:id", s.GetBackUpById)

}
