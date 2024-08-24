package tests

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"

	interfaces "github.com/LanternNassi/IMSController/internal/Interfaces"
	"github.com/LanternNassi/IMSController/internal/server"

	"github.com/LanternNassi/IMSController/internal/database"
	"github.com/LanternNassi/IMSController/internal/models"
	"github.com/stretchr/testify/suite"
)

type ServerSuite struct {
	suite.Suite
	// databaseOperations interfaces.DataBaseClient
	serverOperations interfaces.Server
	server           *echo.Echo
	db               *gorm.DB
	test_client_id   *string
	// test_backup_id       uint
	// test_installation_id uint
	// test_bill_id uint
}

func (s *ServerSuite) SetupSuite() {

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

	new_echo_server, sub_echo := server.NewEchoServer(newdbClient)

	s.serverOperations = new_echo_server
	s.server = sub_echo
	s.db = test_db

}

func (s *ServerSuite) TearDownSuite() {
	//Drop all tables and reset the database
	s.db.Migrator().DropTable(&models.Client{})
	s.db.Migrator().DropTable(&models.Backup{})
	s.db.Migrator().DropTable(&models.Bill{})
	s.db.Migrator().DropTable(&models.Installation{})

}

func NewTestServerSuite() (*ServerSuite, error) {
	return &ServerSuite{}, nil
}
