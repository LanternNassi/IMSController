package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/LanternNassi/IMSController/internal/models"

	interfaces "github.com/LanternNassi/IMSController/internal/Interfaces"
)

type Client struct {
	DB *gorm.DB
}

func NewDatabaseClient() (interfaces.DataBaseClient, error) {

	err_env := godotenv.Load(".env")
	if err_env != nil {
		log.Fatalf("Error loading environment variables file")
	}

	dbport, err_conv := strconv.Atoi(os.Getenv("DBPORT"))

	if err_conv != nil {
		fmt.Println("Error converting string to int:", err_conv)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", os.Getenv("DBHOST"), os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"), os.Getenv("DBNAME"), dbport, "disable")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{

		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		QueryFields: true,
	})

	if err != nil {
		return nil, err
	}

	client := Client{db}

	return client, nil

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
	)

	return err
}
