package tests

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/gorm"

	interfaces "github.com/LanternNassi/IMSController/internal/Interfaces"
	"github.com/LanternNassi/IMSController/internal/database"
	"github.com/LanternNassi/IMSController/internal/models"
	"github.com/stretchr/testify/suite"
)

type DatabaseSuite struct {
	suite.Suite
	databaseOperations interfaces.DataBaseClient
	db                 *gorm.DB

	test_client_id *string
	test_backup_id uint
}

func (s *DatabaseSuite) SetupSuite() {

	//Test Database setup
	err_env := godotenv.Load(".env")
	if err_env != nil {
		log.Fatalf("Error loading environment variables file")
	}
	dbport, err_conv := strconv.Atoi(os.Getenv("test_DBPORT"))
	if err_conv != nil {
		fmt.Println("Error converting string to int:", err_conv)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", os.Getenv("test_DBHOST"), os.Getenv("test_DBUSER"), os.Getenv("test_DBPASSWORD"), os.Getenv("test_DBNAME"), dbport, "disable")

	newdbClient, test_db, err := database.NewDatabaseClient(dsn)

	if err != nil {
		fmt.Print("Error creating database client")
	}

	//Migrating the test database
	migration_err := newdbClient.Migrate()

	if err != nil || migration_err != nil {
		log.Fatalf("Database not loaded ...")

	}

	s.db = test_db
	s.databaseOperations = newdbClient
}

func (s *DatabaseSuite) TearDownSuite() {
	//Drop all tables and reset the database
	s.db.Migrator().DropTable(&models.Client{})
	s.db.Migrator().DropTable(&models.Backup{})
	s.db.Migrator().DropTable(&models.Bill{})
	s.db.Migrator().DropTable(&models.Installation{})

}

func NewTestDatabaseSuite() (*DatabaseSuite, error) {
	return &DatabaseSuite{}, nil
}
