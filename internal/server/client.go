package server

import (
	"net/http"
	"time"

	"github.com/LanternNassi/IMSController/internal/models"
	"github.com/labstack/echo"
)

func (s *EchoServer) AddClient(ctx echo.Context) error {

	client := new(models.Client)

	// type Date struct {
	// 	ValidTill string
	// }

	if err := ctx.Bind(client); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	//Working on the data gotten
	// Date_stamp := new(Date)
	// if date_err := ctx.Bind(Date_stamp); date_err != nil {
	// 	return ctx.JSON(http.StatusBadRequest, date_err)
	// }

	// validDate, date_conv_err := time.Parse(time.RFC3339, Date_stamp.ValidTill)
	// client.ValidTill = validDate

	// if date_conv_err != nil {
	// 	return ctx.JSON(http.StatusBadRequest, date_conv_err)
	// }

	client, err := s.DB.AddClient(ctx.Request().Context(), client)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)

	}

	return ctx.JSON(http.StatusCreated, client)
}

func (s *EchoServer) UpdateClient(ctx echo.Context) error {
	ID := ctx.Param("id")

	type Date struct {
		ValidTill string
	}

	client := new(models.Client)
	if err := ctx.Bind(client); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	active_client, _ := s.DB.GetClientById(ctx.Request().Context(), ID)

	//Working on the data gotten
	Date_stamp := new(Date)
	if date_err := ctx.Bind(Date_stamp); date_err != nil {
		return ctx.JSON(http.StatusBadRequest, date_err)
	}

	if Date_stamp.ValidTill != "" {
		validDate, date_conv_err := time.Parse(time.RFC3339, Date_stamp.ValidTill)

		if date_conv_err != nil {
			return ctx.JSON(http.StatusBadRequest, date_conv_err)
		}

		client.ValidTill = validDate
	} else {
		client.ValidTill = active_client.ValidTill
	}

	//Cleaning up the client object
	if client.FirstName == "" {
		client.FirstName = active_client.FirstName
	}

	if client.LastName == "" {
		client.LastName = active_client.LastName
	}

	if client.Email == "" {
		client.Email = active_client.Email
	}

	if client.Phone == "" {
		client.Phone = active_client.Phone
	}

	if client.Address == "" {
		client.Address = active_client.Address
	}

	if client.BusinessName == "" {
		client.BusinessName = active_client.BusinessName
	}

	if client.Status == "" {
		client.Status = active_client.Status
	}

	client, errs := s.DB.UpdateClient(ctx.Request().Context(), client, ID)

	if errs != nil {
		return ctx.JSON(http.StatusInternalServerError, errs)

	}

	return ctx.JSON(http.StatusOK, client)
}

func (s *EchoServer) GetClientById(ctx echo.Context) error {
	id := ctx.Param("id")

	client, err := s.DB.GetClientById(ctx.Request().Context(), id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, client)
}

func (s *EchoServer) GetClients(ctx echo.Context) error {
	client := new(models.Client)
	if err := ctx.Bind(client); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	clients, errs := s.DB.GetClients(ctx.Request().Context(), client)

	if errs != nil {
		return ctx.JSON(http.StatusInternalServerError, errs)
	}

	return ctx.JSON(http.StatusOK, clients)

}
