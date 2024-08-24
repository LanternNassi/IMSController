package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo"

	"github.com/LanternNassi/IMSController/internal/models"
)

func (s *ServerSuite) Test_001_AddClient() {

	type client struct {
		FirstName    string
		LastName     string
		Email        string
		Phone        string
		Address      string
		BusinessName string
		Status       string
		ValidTill    string
	}

	body, err := json.Marshal(client{
		FirstName:    "TestClient",
		LastName:     "TestClient",
		Email:        "nessim@gmail.com",
		Phone:        "1234567890",
		Address:      "Test Address",
		BusinessName: "Test Business",
		Status:       "Active",
		ValidTill:    "2024-08-22T15:30:00Z",
	})

	s.NoError(err)

	req, err := http.NewRequest(http.MethodPost, "/clients", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	s.NoError(err)

	rec := httptest.NewRecorder()

	c := s.server.NewContext(req, rec)

	err = s.serverOperations.AddClient(c)
	s.NoError(err)
	s.Equal(http.StatusCreated, rec.Code)

	var response models.Client

	err_2 := json.Unmarshal(rec.Body.Bytes(), &response)
	s.NoError(err_2)

	s.test_client_id = &response.ClientID
	s.Equal("TestClient", response.FirstName)

}

func (s *ServerSuite) Test_002_GetClients() {

	req, err := http.NewRequest(http.MethodGet, "/clients", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	s.NoError(err)

	rec := httptest.NewRecorder()

	c := s.server.NewContext(req, rec)

	err_2 := s.serverOperations.GetClients(c)
	s.NoError(err_2)
	s.Equal(http.StatusOK, rec.Code)

	var clients []models.Client

	err_3 := json.Unmarshal(rec.Body.Bytes(), &clients)
	s.NoError(err_3)

	s.Greater(len(clients), 0)

}
