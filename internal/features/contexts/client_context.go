package contexts

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/cucumber/godog"
	"github.com/labstack/echo"

	"github.com/LanternNassi/IMSController/internal/models"
)

type ClientContext struct {
	base           *BaseContext
	client         *models.Client
	response       *httptest.ResponseRecorder
	statusCode     int
	responseBody   map[string]interface{}
	storedClientID string
}

func NewClientContext() *ClientContext {
	return &ClientContext{
		base: GetBaseContext(),
	}
}

func (c *ClientContext) Reset() {
	c.client = nil
	c.response = nil
	c.statusCode = 0
	c.responseBody = nil
	c.storedClientID = ""
}

func (c *ClientContext) SetClient(client *models.Client) {
	c.client = client
}

func (c *ClientContext) GetClient() *models.Client {
	return c.client
}

func (c *ClientContext) SetResponse(response *httptest.ResponseRecorder) {
	c.response = response
}

func (c *ClientContext) GetResponse() *httptest.ResponseRecorder {
	return c.response
}

func (c *ClientContext) SetStatusCode(statusCode int) {
	c.statusCode = statusCode
}

func (c *ClientContext) GetStatusCode() int {
	return c.statusCode
}

func (c *ClientContext) GetResponseBody() map[string]interface{} {
	return c.responseBody
}

func (c *ClientContext) CreateClient(client *models.Client) error {
	if err := c.base.SetupTestDatabase(); err != nil {
		return err
	}

	_, err := c.base.DBClient.AddClient(context.Background(), client)

	if err != nil {
		return fmt.Errorf("error creating client: %w", err)
	}

	return nil
}

func (c *ClientContext) SendCreateClientRequest() error {
	if err := c.base.SetupTestDatabase(); err != nil {
		return err
	}

	if c.client == nil {
		return fmt.Errorf("client is nil, cannot send request")
	}

	if c.base.Server == nil {
		return fmt.Errorf("server is not initialized")
	}

	// Format date as RFC3339 (required by Echo's JSON binding)
	validTillStr := c.client.ValidTill.Format(time.RFC3339)

	clientData := map[string]interface{}{
		"FirstName":    c.client.FirstName,
		"LastName":     c.client.LastName,
		"Email":        c.client.Email,
		"Phone":        c.client.Phone,
		"Address":      c.client.Address,
		"BusinessName": c.client.BusinessName,
		"Status":       c.client.Status,
		"ValidTill":    validTillStr,
	}

	jsonData, err := json.Marshal(clientData)
	if err != nil {
		return fmt.Errorf("error marshalling client data: %w", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/clients", bytes.NewBuffer(jsonData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c.base.Server.ServeHTTP(rec, req)
	c.response = rec
	c.statusCode = rec.Code

	// Parse response body
	if rec.Body.Len() > 0 {
		var body map[string]interface{}
		if err := json.Unmarshal(rec.Body.Bytes(), &body); err == nil {
			c.responseBody = body
		} else {
			fmt.Printf("Warning: Could not parse response body as JSON: %v. Body: %s\n", err, rec.Body.String())
		}
	} else {
		fmt.Printf("Warning: Empty response body. Status: %d\n", rec.Code)
	}

	return nil
}

func (c *ClientContext) Base() *BaseContext {
	return c.base
}

func (c *ClientContext) StoreClientIDFromResponse() error {
	if c.responseBody == nil {
		return fmt.Errorf("response body is nil")
	}

	clientID, ok := c.responseBody["ClientID"]
	if !ok {
		return fmt.Errorf("response does not contain ClientID")
	}

	clientIDStr, ok := clientID.(string)
	if !ok {
		return fmt.Errorf("ClientID is not a string")
	}

	c.storedClientID = clientIDStr
	return nil
}

func (c *ClientContext) GetStoredClientID() string {
	return c.storedClientID
}

func (c *ClientContext) SendGetClientRequest(clientID string) error {
	if err := c.base.SetupTestDatabase(); err != nil {
		return err
	}

	if c.base.Server == nil {
		return fmt.Errorf("server is not initialized")
	}

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/clients/%s", clientID), nil)
	rec := httptest.NewRecorder()

	c.base.Server.ServeHTTP(rec, req)
	c.response = rec
	c.statusCode = rec.Code

	if rec.Body.Len() > 0 {
		var body map[string]interface{}
		if err := json.Unmarshal(rec.Body.Bytes(), &body); err == nil {
			c.responseBody = body
		} else {
			fmt.Printf("Warning: Could not parse response body as JSON: %v. Body: %s\n", err, rec.Body.String())
		}
	}

	return nil
}

func (c *ClientContext) SendGetClientsRequest() error {
	if err := c.base.SetupTestDatabase(); err != nil {
		return err
	}

	if c.base.Server == nil {
		return fmt.Errorf("server is not initialized")
	}

	req := httptest.NewRequest(http.MethodGet, "/clients", nil)
	rec := httptest.NewRecorder()

	c.base.Server.ServeHTTP(rec, req)
	c.response = rec
	c.statusCode = rec.Code

	if rec.Body.Len() > 0 {
		var body []interface{}
		if err := json.Unmarshal(rec.Body.Bytes(), &body); err == nil {
			c.responseBody = map[string]interface{}{"clients": body}
		} else {
			// Try as single object if array fails
			var singleBody map[string]interface{}
			if err2 := json.Unmarshal(rec.Body.Bytes(), &singleBody); err2 == nil {
				c.responseBody = singleBody
			} else {
				fmt.Printf("Warning: Could not parse response body: %v. Body: %s\n", err2, rec.Body.String())
			}
		}
	}

	return nil
}

// SendUpdateClientRequest sends a PUT request to update a client
func (c *ClientContext) SendUpdateClientRequest(clientID string, updates map[string]interface{}) error {
	if err := c.base.SetupTestDatabase(); err != nil {
		return err
	}

	if c.base.Server == nil {
		return fmt.Errorf("server is not initialized")
	}

	if updates == nil {
		updates = make(map[string]interface{})
	}

	jsonData, err := json.Marshal(updates)
	if err != nil {
		return fmt.Errorf("error marshalling update data: %w", err)
	}

	// Create request with proper body
	bodyBuffer := bytes.NewBuffer(jsonData)
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/clients/%s", clientID), bodyBuffer)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.ContentLength = int64(len(jsonData))

	rec := httptest.NewRecorder()

	c.base.Server.ServeHTTP(rec, req)
	c.response = rec
	c.statusCode = rec.Code

	if rec.Body.Len() > 0 {
		var body map[string]interface{}
		if err := json.Unmarshal(rec.Body.Bytes(), &body); err == nil {
			c.responseBody = body
		} else {
			fmt.Printf("Warning: Could not parse update response body as JSON: %v. Body: %s\n", err, rec.Body.String())
		}
	} else {
		fmt.Printf("Warning: Empty response body for update. Status: %d\n", rec.Code)
	}

	return nil
}

func ParseClientFromTable(table *godog.Table) (*models.Client, error) {
	if len(table.Rows) < 2 {
		return nil, fmt.Errorf("table must have at least a header and one data row")
	}

	headers := table.Rows[0].Cells
	dataRow := table.Rows[1].Cells

	client := &models.Client{}
	validTill := time.Now().AddDate(1, 0, 0) // Default to 1 year from now

	for i, header := range headers {
		if i >= len(dataRow) {
			continue
		}

		value := dataRow[i].Value
		headerValue := header.Value

		switch headerValue {
		case "FirstName":
			client.FirstName = value
		case "LastName":
			client.LastName = value
		case "Email":
			client.Email = value
		case "Phone":
			client.Phone = value
		case "Address":
			client.Address = value
		case "BusinessName":
			client.BusinessName = value
		case "Status":
			client.Status = value
		case "ValidTill":
			if value != "" {
				parsedTime, err := time.Parse("2006-01-02", value)
				if err != nil {
					return nil, fmt.Errorf("invalid date format for ValidTill: %w", err)
				}
				validTill = parsedTime.UTC()
			}
		}
	}

	client.ValidTill = validTill
	return client, nil
}
