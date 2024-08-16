package tests

import (
	"context"
	"time"

	"github.com/LanternNassi/IMSController/internal/models"
)

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
	s.test_client_id = &returned_client.ClientID
	s.NoError(err)
}

func (s *DatabaseSuite) TestGetClients() {
	clients, err := s.databaseOperations.GetClients(context.Background(), &models.Client{})
	s.NoError(err)
	s.Equal(1, len(clients))
}

func (s *DatabaseSuite) TestGetClientById() {
	client, err := s.databaseOperations.GetClientById(context.Background(), *s.test_client_id)
	s.NoError(err)
	s.Equal("TestClient", client.FirstName)
}

func (s *DatabaseSuite) TestUpdateClient() {
	updated_client := models.Client{
		FirstName: "UpdatedClient",
	}

	_, err := s.databaseOperations.UpdateClient(context.Background(), &updated_client, *s.test_client_id)

	s.NoError(err)

	client, get_err := s.databaseOperations.GetClientById(context.Background(), *s.test_client_id)

	s.NoError(get_err)
	s.Equal("UpdatedClient", client.FirstName)
}
