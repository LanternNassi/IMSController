package steps

import (
	"fmt"

	"github.com/LanternNassi/IMSController/internal/features/contexts"
	"github.com/cucumber/godog"
)

var clientCtx *contexts.ClientContext

func GetClientContext() *contexts.ClientContext {
	return clientCtx
}

// InitializeClientSteps registers all client-related step definitions
func InitializeClientSteps(ctx *godog.ScenarioContext) {
	clientCtx = contexts.NewClientContext()

	ctx.Step(`^the system has a test database$`, theSystemHasATestDatabaseForClient)
	ctx.Step(`^a client exists with the following details:$`, aClientExistsWithTheFollowingDetails)
	ctx.Step(`^I send a request to create a client with the following details:$`, iSendARequestToCreateAClientWithTheFollowingDetails)
	ctx.Step(`^I send a request to get client with ID "([^"]*)"$`, iSendARequestToGetClientWithID)
	ctx.Step(`^I send a request to get client with ID from the response$`, iSendARequestToGetClientWithIDFromResponse)
	ctx.Step(`^I send a request to get all clients$`, iSendARequestToGetAllClients)
	ctx.Step(`^I send a request to update client with ID "([^"]*)" with the following details:$`, iSendARequestToUpdateClientWithID)
	ctx.Step(`^I send a request to update client with ID from stored ID with the following details:$`, iSendARequestToUpdateClientWithIDFromStored)
	ctx.Step(`^I store the client ID from the response$`, iStoreTheClientIDFromTheResponse)
	ctx.Step(`^the response should contain client ID$`, theResponseShouldContainClientID)
	ctx.Step(`^the response should contain FirstName "([^"]*)"$`, theResponseShouldContainFirstName)
	ctx.Step(`^the response should contain LastName "([^"]*)"$`, theResponseShouldContainLastName)
	ctx.Step(`^the response should contain Email "([^"]*)"$`, theResponseShouldContainEmailForClient)
}

func theSystemHasATestDatabaseForClient() error {
	return clientCtx.Base().SetupTestDatabase()
}

func aClientExistsWithTheFollowingDetails(table *godog.Table) error {
	client, err := contexts.ParseClientFromTable(table)
	if err != nil {
		return fmt.Errorf("error parsing client from table: %w", err)
	}
	clientCtx.SetClient(client)
	return clientCtx.SendCreateClientRequest()
}

func iSendARequestToCreateAClientWithTheFollowingDetails(table *godog.Table) error {
	client, err := contexts.ParseClientFromTable(table)
	if err != nil {
		return fmt.Errorf("error parsing client from table: %w", err)
	}

	clientCtx.SetClient(client)
	return clientCtx.SendCreateClientRequest()
}

func iSendARequestToGetClientWithID(clientID string) error {
	return clientCtx.SendGetClientRequest(clientID)
}

func iSendARequestToGetAllClients() error {
	return clientCtx.SendGetClientsRequest()
}

func iSendARequestToUpdateClientWithID(clientID string, table *godog.Table) error {
	if len(table.Rows) < 2 {
		return fmt.Errorf("table must have at least a header and one data row")
	}

	headers := table.Rows[0].Cells
	dataRow := table.Rows[1].Cells
	updates := make(map[string]interface{})

	for i, header := range headers {
		if i >= len(dataRow) {
			continue
		}
		updates[header.Value] = dataRow[i].Value
	}

	return clientCtx.SendUpdateClientRequest(clientID, updates)
}

func theResponseShouldContainClientID() error {
	responseBody := clientCtx.GetResponseBody()
	if responseBody == nil {
		return fmt.Errorf("response body is nil")
	}

	clientID, ok := responseBody["ClientID"]
	if !ok || clientID == nil {
		return fmt.Errorf("response does not contain ClientID. Response: %v", responseBody)
	}

	clientIDStr, ok := clientID.(string)
	if !ok || clientIDStr == "" {
		return fmt.Errorf("ClientID is not a valid string. Response: %v", responseBody)
	}

	return nil
}

func theResponseShouldContainFirstName(firstName string) error {
	responseBody := clientCtx.GetResponseBody()
	if responseBody == nil {
		return fmt.Errorf("response body is nil")
	}

	responseFirstName, ok := responseBody["FirstName"]
	if !ok {
		return fmt.Errorf("response does not contain FirstName. Response: %v", responseBody)
	}

	firstNameStr, ok := responseFirstName.(string)
	if !ok || firstNameStr != firstName {
		return fmt.Errorf("expected FirstName %s, but got %v. Response: %v", firstName, responseFirstName, responseBody)
	}

	return nil
}

func iSendARequestToGetClientWithIDFromResponse() error {
	if err := clientCtx.StoreClientIDFromResponse(); err != nil {
		return err
	}
	return clientCtx.SendGetClientRequest(clientCtx.GetStoredClientID())
}

func iSendARequestToUpdateClientWithIDFromStored(table *godog.Table) error {
	clientID := clientCtx.GetStoredClientID()
	if clientID == "" {
		return fmt.Errorf("no stored client ID available")
	}

	if len(table.Rows) < 2 {
		return fmt.Errorf("table must have at least a header and one data row")
	}

	headers := table.Rows[0].Cells
	dataRow := table.Rows[1].Cells
	updates := make(map[string]interface{})

	for i, header := range headers {
		if i >= len(dataRow) {
			continue
		}
		updates[header.Value] = dataRow[i].Value
	}

	
	return clientCtx.SendUpdateClientRequest(clientID, updates)
}

func iStoreTheClientIDFromTheResponse() error {
	return clientCtx.StoreClientIDFromResponse()
}

func theResponseShouldContainEmailForClient(email string) error {
	responseBody := clientCtx.GetResponseBody()
	if responseBody == nil {
		return fmt.Errorf("response body is nil")
	}

	responseEmail, ok := responseBody["Email"]
	if !ok {
		return fmt.Errorf("response does not contain Email. Response: %v", responseBody)
	}

	emailStr, ok := responseEmail.(string)
	if !ok || emailStr != email {
		return fmt.Errorf("expected Email %s, but got %v. Response: %v", email, responseEmail, responseBody)
	}

	return nil
}

func theResponseShouldContainLastName(lastName string) error {
	responseBody := clientCtx.GetResponseBody()
	if responseBody == nil {
		return fmt.Errorf("response body is nil")
	}

	responseLastName, ok := responseBody["LastName"]
	if !ok {
		return fmt.Errorf("response does not contain LastName. Response: %v", responseBody)
	}

	lastNameStr, ok := responseLastName.(string)
	if !ok || lastNameStr != lastName {
		return fmt.Errorf("expected LastName %s, but got %v. Response: %v", lastName, responseLastName, responseBody)
	}

	return nil
}