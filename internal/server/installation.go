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

func (s *EchoServer) UpdateInstallation(ctx echo.Context) error {

	id := ctx.Param("id")
	inst, inst_err := s.DB.GetInstallationById(ctx.Request().Context(), id)

	if inst_err != nil {
		return ctx.JSON(http.StatusNotFound, inst_err)
	}

	installation := new(models.Installation)

	if err := ctx.Bind(installation); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	//Object cleaning up
	if installation.Installation_type == "" {
		installation.Installation_type = inst.Installation_type
	}

	if installation.Computer_name == "" {
		installation.Computer_name = inst.Computer_name
	}

	if installation.IMS_version == "" {
		installation.IMS_version = inst.IMS_version
	}

	if installation.Operating_system == "" {
		installation.Operating_system = inst.Operating_system
	}

	if installation.RAM == "" {
		installation.RAM = inst.RAM
	}

	if installation.Processor == "" {
		installation.Processor = inst.Processor
	}

	if installation.Active == "" {
		installation.Active = inst.Active
	}

	returned_installation, err_2 := s.DB.UpdateInstallation(ctx.Request().Context(), installation, id)

	if err_2 != nil {
		return ctx.JSON(http.StatusInternalServerError, err_2)
	}

	return ctx.JSON(http.StatusOK, returned_installation)
}
