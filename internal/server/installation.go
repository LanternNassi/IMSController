package server

import (
	"net/http"

	"github.com/LanternNassi/IMSController/internal/models"
	"github.com/labstack/echo"
)

func (s *EchoServer) GetInstallations(ctx echo.Context) error {
	installation := new(models.Installation)

	if err := ctx.Bind(installation); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)

	}

	installations, err := s.DB.GetInstallations(ctx.Request().Context(), installation)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, installations)
}

func (s *EchoServer) GetInstallationById(ctx echo.Context) error {
	id := ctx.Param("id")

	installation, err := s.DB.GetInstallationById(ctx.Request().Context(), id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, installation)
}

func (s *EchoServer) AddInstallation(ctx echo.Context) error {
	installation := new(models.Installation)
	if err := ctx.Bind(installation); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	installation, err := s.DB.AddInstallation(ctx.Request().Context(), installation)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusCreated, installation)
}
