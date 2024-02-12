package server

import (
	"net/http"

	"github.com/LanternNassi/IMSController/internal/models"
	"github.com/labstack/echo"
)

func (s *EchoServer) GetBills(ctx echo.Context) error {
	bill := new(models.Bill)

	if err := ctx.Bind(bill); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	bills, err := s.DB.GetBills(ctx.Request().Context(), bill)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, bills)
}

func (s *EchoServer) AddBill(ctx echo.Context) error {

	bill := new(models.Bill)
	if err := ctx.Bind(bill); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	bill, err := s.DB.AddBill(ctx.Request().Context(), bill)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)

	}

	return ctx.JSON(http.StatusCreated, bill)
}

func (s *EchoServer) GetBillById(ctx echo.Context) error {
	id := ctx.Param("id")

	bill, err := s.DB.GetBillById(ctx.Request().Context(), id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, bill)
}
