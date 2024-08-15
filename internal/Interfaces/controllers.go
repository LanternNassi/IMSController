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
	DeleteBackUpById(ctx context.Context, id string) (bool, error)

	GetBills(ctx context.Context, params *models.Bill) ([]models.Bill, error)
	AddBill(ctx context.Context, bill *models.Bill) (*models.Bill, error)
	UpdateBill(ctx context.Context, bill *models.Bill, id string) (*models.Bill, error)
	GetBillById(ctx context.Context, id string) (*models.Bill, error)
	GetBillsByDate(ctx context.Context, field string, comparator string, time_var time.Time, client_id string) ([]models.Bill, error)

	GetInstallations(ctx context.Context, params *models.Installation) ([]models.Installation, error)
	AddInstallation(ctx context.Context, installation *models.Installation) (*models.Installation, error)
	UpdateInstallation(ctx context.Context, installation *models.Installation, id string) (*models.Installation, error)
	GetInstallationById(ctx context.Context, id string) (*models.Installation, error)
	DeleteInstallation(ctx context.Context, is string) (bool, error)
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
	GetBackUpByClientId(ctx echo.Context) error
	GetBackUpByBill(ctx echo.Context) error
	DeleteBackUpById(ctx echo.Context) error

	GetBills(ctx echo.Context) error
	AddBill(ctx echo.Context) error
	GetBillById(ctx echo.Context) error
	UpdateBill(ctx echo.Context) error
	GetBillByClientId(ctx echo.Context) error

	GetInstallations(ctx echo.Context) error
	AddInstallation(ctx echo.Context) error
	GetInstallationById(ctx echo.Context) error
}
