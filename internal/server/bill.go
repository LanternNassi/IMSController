package server

import (
	"net/http"
	"time"

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

func (s *EchoServer) UpdateBill(ctx echo.Context) error {
	ID := ctx.Param("id")

	_bill := new(models.Bill)
	if err := ctx.Bind(_bill); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	active_bill, _ := s.DB.GetBillById(ctx.Request().Context(), ID)

	if _bill.Billed {
		//Clearing the clients disconnected state and prolonging their activity
		client, client_errs := s.DB.GetClientById(ctx.Request().Context(), active_bill.ClientID)

		if client_errs != nil {
			return ctx.JSON(http.StatusBadRequest, client_errs)
		}

		client.Status = "connected"
		client.ValidTill = time.Now().AddDate(0, 0, 30)

		_, _client_errs := s.DB.UpdateClient(ctx.Request().Context(), client, active_bill.ClientID)

		if _client_errs != nil {
			return ctx.JSON(http.StatusBadRequest, _client_errs)
		}

	}

	bill, errs := s.DB.UpdateBill(ctx.Request().Context(), _bill, ID)

	if errs != nil {
		return ctx.JSON(http.StatusInternalServerError, errs)
	}

	return ctx.JSON(http.StatusOK, bill)

}
