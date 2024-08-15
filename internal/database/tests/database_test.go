package tests

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/gorm"

	interfaces "github.com/LanternNassi/IMSController/internal/Interfaces"
	databaseops "github.com/LanternNassi/IMSController/internal/database"
	"github.com/LanternNassi/IMSController/internal/models"
	"github.com/stretchr/testify/suite"
)

type DatabaseSuite struct {
	suite.Suite
	databaseOperations interfaces.DataBaseClient
	db                 *gorm.DB

	test_client_id string
}

func (s *DatabaseSuite) SetupSuite() {

	//Test Database setup
	err_env := godotenv.Load(".test_env")
	if err_env != nil {
		log.Fatalf("Error loading environment variables file")
	}
	dbport, err_conv := strconv.Atoi(os.Getenv("test_DBPORT"))
	if err_conv != nil {
		fmt.Println("Error converting string to int:", err_conv)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", os.Getenv("test_DBHOST"), os.Getenv("test_DBUSER"), os.Getenv("test_DBPASSWORD"), os.Getenv("test_DBNAME"), dbport, "disable")

	newdbClient, test_db, err := databaseops.NewDatabaseClient(dsn)

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

func (s *DatabaseSuite) TestAddClient() {
	client := models.Client{
		FirstName:    "TestClient",
		LastName:     "TestClient",
		Email:        "nessim@gmail.com",
		Phone:        "1234567890",
		Address:      "Test Address",
		BusinessName: "Test Business",
		Status:       "Active",
		ValidTill:    time.Date(2022, 12, 12, 0, 0, 0, 0, time.UTC),
	}
	returned_client, err := s.databaseOperations.AddClient(context.Background(), &client)
	s.test_client_id = returned_client.ClientID
	s.NoError(err)
}

func (s *DatabaseSuite) TestGetClients() {
	clients, err := s.databaseOperations.GetClients(context.Background(), &models.Client{})
	s.NoError(err)
	s.Equal(1, len(clients))
}

func (s *DatabaseSuite) TestGetClientById() {
	client, err := s.databaseOperations.GetClientById(context.Background(), s.test_client_id)
	s.NoError(err)
	s.Equal("TestClient", client.FirstName)
}

func (s *DatabaseSuite) TestUpdateClient() {
	updated_client := models.Client{
		FirstName: "UpdatedClient",
	}

	_, err := s.databaseOperations.UpdateClient(context.Background(), &updated_client, s.test_client_id)

	s.NoError(err)

	client, get_err := s.databaseOperations.GetClientById(context.Background(), s.test_client_id)

	s.NoError(get_err)
	s.Equal("UpdatedClient", client.FirstName)
}

func TestDatabaseSuite(t *testing.T) {
	suite.Run(t, new(DatabaseSuite))
}
