package server

import (
	"net/http"

	"time"

	"strconv"

	"github.com/LanternNassi/IMSController/internal/models"
	"github.com/labstack/echo"

	"github.com/shopspring/decimal"

	"io"
	"mime/multipart"
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

	//Handling file upload
	max_limit_err := ctx.Request().ParseMultipartForm(10 << 20)

	if max_limit_err != nil {
		return ctx.JSON(http.StatusBadRequest, max_limit_err)
	}

	//Retrieving the file
	file, handler, file_err := ctx.Request().FormFile("file")
	if file_err != nil {
		return ctx.JSON(http.StatusBadRequest, file_err)
	}

	defer func(file multipart.File) {
		file.Close()

	}(file)

	// Read the content of the file
	fileBytes, content_err := io.ReadAll(file)

	if content_err != nil {
		return ctx.JSON(http.StatusBadRequest, content_err)
	}

	backup := new(models.Backup)
	if err := ctx.Bind(backup); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	//Determining the bill to add the backup to ...

	var _bill *models.Bill

	if backup.Bill == 0 {

		//Creating a new  bill or determining the bill to append the backup

		currentTime := time.Now()

		year, month, _ := currentTime.Date()

		firstDayOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, currentTime.Location())

		bills, bill_err := s.DB.GetBillsByDate(ctx.Request().Context(), "created_at", ">", firstDayOfMonth, backup.ClientID)

		if bill_err != nil {
			return ctx.JSON(http.StatusBadRequest, bill_err)
		}

		if len(bills) <= 0 {

			_bill, _ = s.DB.AddBill(ctx.Request().Context(), &models.Bill{
				ClientID: backup.ClientID,
			})

		} else {

			_bill = &bills[0]

		}

	} else {
		bill, bill_err := s.DB.GetBillById(ctx.Request().Context(), strconv.FormatUint(uint64(backup.Bill), 10))

		if bill_err != nil {
			return ctx.JSON(http.StatusBadGateway, bill_err)
		}

		_bill = bill

	}

	//Performing operations on the billing object
	_bill.BackupCount += 1

	//Adding to the bill backup size
	_bill.BackupSize += handler.Size

	//Adding the cost based on the size (each byte costing 0.001 to 0.002 UGx)
	_bill.TotalCost = _bill.TotalCost.Add(decimal.NewFromFloat(float64(handler.Size) * 0.0018273998877))

	//Updating the bill
	_bill, _ = s.DB.UpdateBill(ctx.Request().Context(), _bill, strconv.FormatUint(uint64(_bill.ID), 10))

	// Adding the file specifications to the model
	backup.Name = handler.Filename
	backup.Size = handler.Size
	backup.Backup = fileBytes
	backup.Bill = _bill.ID

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

func (s *EchoServer) DeleteBackUpById(ctx echo.Context) error {
	id := ctx.Param("id")

	deleted, err := s.DB.DeleteBackUpById(ctx.Request().Context(), id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, deleted)
}

func (s *EchoServer) GetBackUpByClientId(ctx echo.Context) error {
	clientId := ctx.Param("Id")
	backups, err := s.DB.Getbackups(ctx.Request().Context(), &models.Backup{ClientID: clientId})

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, backups)
}

func (s *EchoServer) GetBackUpByBill(ctx echo.Context) error {
	bill := ctx.Param("bill")

	conv_bill, err := strconv.ParseUint(bill, 10, 32)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	backups, err := s.DB.Getbackups(ctx.Request().Context(), &models.Backup{Bill: uint(conv_bill)})

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, backups)
}

func (s *EchoServer) DownloadBackup(ctx echo.Context) error {
	id := ctx.Param("id")

	backup, err := s.DB.GetBackUpById(ctx.Request().Context(), id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.Response().Header().Set("Content-Disposition", "attachment; filename="+backup.Name)

	return ctx.Blob(http.StatusOK, "application/octet-stream", backup.Backup)

}
