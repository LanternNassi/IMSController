package server

import (
	"net/http"

	"github.com/LanternNassi/IMSController/internal/models"
	"github.com/labstack/echo"
)

func (s *EchoServer) Getbackups(ctx echo.Context) error {
	backup := new(models.Backup)

	if err := ctx.Bind(backup); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	backups, err := s.DB.Getbackups(ctx.Request().Context(), backup)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, backups)

}

func (s *EchoServer) AddBackup(ctx echo.Context) error {
	backup := new(models.Backup)
	if err := ctx.Bind(backup); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	backup, err := s.DB.AddBackup(ctx.Request().Context(), backup)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)

	}

	return ctx.JSON(http.StatusCreated, backup)

}

func (s *EchoServer) GetBackUpById(ctx echo.Context) error {
	id := ctx.Param("id")

	backup, err := s.DB.GetBackUpById(ctx.Request().Context(), id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, backup)
}
