package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	app_middleware "github.com/LanternNassi/IMSController/internal/middleware"
	"github.com/LanternNassi/IMSController/internal/models"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"

	interfaces "github.com/LanternNassi/IMSController/internal/Interfaces"
)

type EchoServer struct {
	echo *echo.Echo
	DB   interfaces.DataBaseClient
}

func NewEchoServer(db interfaces.DataBaseClient) (interfaces.Server, *echo.Echo) {
	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}

	server.registerRoutes()
	return server, server.echo

}

func (s *EchoServer) Start() error {

	err_env := godotenv.Load(".env")
	if err_env != nil {
		fmt.Println("Error loading environment variables file... Proceeding to use default values")
	}

	if err := s.echo.Start(os.Getenv("APPHOST")); err != nil && err != http.ErrServerClosed {
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

	s.echo.Use(middleware.CORS())

	s.echo.GET("/readiness", s.Readiness)
	s.echo.GET("/liveness", s.Liveness)

	cg := s.echo.Group("/clients")
	cg.GET("", s.GetClients)
	cg.POST("", s.AddClient)
	cg.PUT("/:id", s.UpdateClient)
	cg.GET("/:id", s.GetClientById)

	bg := s.echo.Group("/backups")
	bg.GET("", s.Getbackups, app_middleware.AuthenticationMiddleware())
	bg.POST("", s.AddBackup, app_middleware.AuthenticationMiddleware())
	bg.GET("/:id", s.GetBackUpById, app_middleware.AuthenticationMiddleware())
	bg.GET("/client/:Id", s.GetBackUpByClientId, app_middleware.AuthenticationMiddleware())
	bg.GET("/bill/:bill", s.GetBackUpByBill, app_middleware.AuthenticationMiddleware())
	bg.DELETE("/delete/:id", s.DeleteBackUpById, app_middleware.AuthenticationMiddleware())

	dg := s.echo.Group("/bills")
	dg.GET("", s.GetBills, app_middleware.AuthenticationMiddleware())
	dg.POST("", s.AddBill, app_middleware.AuthenticationMiddleware())
	dg.GET("/:id", s.GetBillById, app_middleware.AuthenticationMiddleware())
	dg.PUT("/:id", s.UpdateBill, app_middleware.AuthenticationMiddleware())
	dg.GET("/client/:ClientId", s.GetBillByClientId, app_middleware.AuthenticationMiddleware())

	Ig := s.echo.Group("/Installations")
	Ig.GET("", s.GetInstallations)
	Ig.POST("", s.AddInstallation)
	Ig.GET("/:id", s.GetInstallationById)

	Ug := s.echo.Group("/Users")
	Ug.GET("", s.GetUsers)
	Ug.POST("", s.AddUser)
	Ug.POST("/login", s.Login)
	Ug.DELETE("/:id", s.DeleteUser)
	Ug.GET("/:id", s.GetUserById)
}
