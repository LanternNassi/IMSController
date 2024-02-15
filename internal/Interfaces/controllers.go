package interfaces

import (
	"context"

	"time"

	"github.com/labstack/echo"

	"github.com/LanternNassi/IMSController/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
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
	GetBackUpsByDate(ctx context.Context, field string, comparator string, time_var time.Time) ([]models.Backup, error)

	GetBills(ctx context.Context, params *models.Bill) ([]models.Bill, error)
	AddBill(ctx context.Context, bill *models.Bill) (*models.Bill, error)
	UpdateBill(ctx context.Context, bill *models.Bill, id string) (*models.Bill, error)
	GetBillById(ctx context.Context, id string) (*models.Bill, error)
	GetBillsByDate(ctx context.Context, field string, comparator string, time_var time.Time , client_id string) ([]models.Bill, error)
}

type MongoDatabaseClient interface {
	CloseMongo(client *mongo.Client, ctx context.Context)
	ConnectMongo(uri string) (*mongo.Client, context.Context, error)
	PingMongo(client *mongo.Client, ctx context.Context) error
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

	GetBills(ctx echo.Context) error
	AddBill(ctx echo.Context) error
	GetBillById(ctx echo.Context) error
	UpdateBill(ctx echo.Context) error
}
