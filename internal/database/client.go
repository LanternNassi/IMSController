package database

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/LanternNassi/IMSController/internal/models"

	interfaces "github.com/LanternNassi/IMSController/internal/Interfaces"
)

type Client struct {
	DB *gorm.DB
}

func NewDatabaseClient(dsn string) (interfaces.DataBaseClient, *gorm.DB, error) {

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", os.Getenv("DBHOST"), os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"), os.Getenv("DBNAME"), dbport, "disable")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{

		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		QueryFields: true,
	})

	if err != nil {
		return nil, nil, err
	}

	client := Client{db}

	return client, client.DB, nil

}

func (client Client) Ready() bool {
	var ready string

	tx := client.DB.Raw("SELECT 1 as ready").Scan(&ready)

	if tx.Error != nil {
		return false
	}

	if ready == "!" {
		return true
	}

	return false
}

func (client Client) Migrate() error {

	err := client.DB.AutoMigrate(
		&models.Client{},
		&models.Backup{},
		&models.Bill{},
		&models.Installation{},
	)

	return err
}
