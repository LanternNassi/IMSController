package contexts

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"gorm.io/gorm"

	interfaces "github.com/LanternNassi/IMSController/internal/Interfaces"
	"github.com/LanternNassi/IMSController/internal/database"
	"github.com/LanternNassi/IMSController/internal/models"
	"github.com/LanternNassi/IMSController/internal/server"
)

// BaseContext provides shared database and server setup for all feature tests
type BaseContext struct {
	Server   *echo.Echo
	DB       *gorm.DB
	DBClient interfaces.DataBaseClient
}

var baseCtx *BaseContext

// GetBaseContext returns the singleton base context instance
func GetBaseContext() *BaseContext {
	if baseCtx == nil {
		baseCtx = &BaseContext{}
	}
	return baseCtx
}

// SetupTestDatabase initializes the test database and server
func (b *BaseContext) SetupTestDatabase() error {
	if b.DB != nil {
		return nil // Already set up
	}

	// Try loading .env from multiple locations
	envPaths := []string{".env", "../../../.env", "../../.env"}
	var errEnv error
	for _, path := range envPaths {
		errEnv = godotenv.Load(path)
		if errEnv == nil {
			fmt.Printf("Loaded environment variables from: %s\n", path)
			break
		}
	}
	if errEnv != nil {
		fmt.Printf("Error loading environment variables file: %v\nProceeding to use default values\n", errEnv)
	}

	dbportStr := os.Getenv("test_DBPORT")
	if dbportStr == "" {
		return fmt.Errorf("test_DBPORT environment variable is not set")
	}

	dbport, errConv := strconv.Atoi(dbportStr)
	if errConv != nil {
		return fmt.Errorf("error converting string to int: %w", errConv)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		os.Getenv("test_DBHOST"),
		os.Getenv("test_DBUSER"),
		os.Getenv("test_DBPASSWORD"),
		os.Getenv("test_DBNAME"),
		dbport,
		"disable")

	newdbClient, testDb, err := database.NewDatabaseClient(dsn)
	if err != nil {
		return fmt.Errorf("error creating database client: %w", err)
	}

	migrationErr := newdbClient.Migrate()
	if migrationErr != nil {
		return fmt.Errorf("error migrating database: %w", migrationErr)
	}


	b.DB = testDb
	b.DBClient = newdbClient

	newEchoServer, echoServer := server.NewEchoServer(newdbClient)
	_ = newEchoServer // Keep reference
	b.Server = echoServer

	return nil
}

// Cleanup drops all test tables and recreates them for the next scenario
func (b *BaseContext) Cleanup() {
	if b.DB != nil && b.DBClient != nil {
		// Drop all tables
		b.DB.Migrator().DropTable(&models.User{})
		b.DB.Migrator().DropTable(&models.Client{})
		b.DB.Migrator().DropTable(&models.Backup{})
		b.DB.Migrator().DropTable(&models.Bill{})
		b.DB.Migrator().DropTable(&models.Installation{})

		// Recreate tables for the next scenario
		// This ensures the next scenario has a fresh database
		if err := b.DBClient.Migrate(); err != nil {
			fmt.Printf("Warning: Failed to recreate tables after cleanup: %v\n", err)
		}
	}
}

// Reset reinitializes the base context (useful for test isolation)
func (b *BaseContext) Reset() {
	b.Server = nil
	b.DB = nil
	b.DBClient = nil
}
