package server

import (
	"net/http"

	"github.com/LanternNassi/IMSController/internal/models"
	"github.com/labstack/echo"

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
		err := file.Close()
		if err != nil {

		}
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

	// Adding the file specifications to the model
	backup.Name = handler.Filename
	backup.Size = handler.Size
	backup.Backup = fileBytes

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

func (s *EchoServer) DownloadBackup(ctx echo.Context) error {
	id := ctx.Param("id")

	backup, err := s.DB.GetBackUpById(ctx.Request().Context(), id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.Response().Header().Set("Content-Disposition", "attachment; filename="+backup.Name)

	return ctx.Blob(http.StatusOK, "application/octet-stream", backup.Backup)

}
