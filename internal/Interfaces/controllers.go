package interfaces

import (
	"context"

	"github.com/labstack/echo"

	"github.com/LanternNassi/IMSController/internal/models"
)

// Interface for creating a Database client
type DataBaseClient interface {
	Ready() bool
	Migrate() error

	GetClients(ctx context.Context, params *models.Client) ([]models.Client, error)
	AddClient(ctx context.Context, client *models.Client) (*models.Client, error)
	UpdateClient(ctx context.Context, client *models.Client, id string) (*models.Client, error)
	GetClientById(ctx context.Context, id string) (*models.Client, error)

	Getbackups(ctx context.Context, params *models.Backup) ([]models.Backup, error)
	AddBackup(ctx context.Context, backup *models.Backup) (*models.Backup, error)
	GetBackUpById(ctx context.Context, id string) (*models.Backup, error)
}

// Interface for creating a Server
type Server interface {
	Start() error
	Readiness(ctx echo.Context) error
	Liveness(ctx echo.Context) error

	GetClients(ctx echo.Context) error
	AddClient(ctx echo.Context) error
	UpdateClient(ctx echo.Context) error
	GetClientById(ctx echo.Context) error

	Getbackups(ctx echo.Context) error
	AddBackup(ctx echo.Context) error
	GetBackUpById(ctx echo.Context) error
}
