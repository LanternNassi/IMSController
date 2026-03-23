package steps

import (
	"bytes"
	"fmt"
	"net/http/httptest"

	"github.com/cucumber/godog"
)

// InitializeCommonSteps registers common step definitions used across multiple features
func InitializeCommonSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^the response status should be (\d+)$`, theResponseStatusShouldBe)
	ctx.Step(`^the response should contain a token$`, theResponseShouldContainAToken)
	ctx.Step(`^the response should contain username$`, theResponseShouldContainUsername)
	ctx.Step(`^the response should contain email "([^"]*)"$`, theResponseShouldContainEmail)
	ctx.Step(`^the response should contain "([^"]*)"$`, theResponseShouldContain)
}

// theResponseStatusShouldBe validates the HTTP status code
func theResponseStatusShouldBe(statusCode int) error {
	var actualCode int
	var response *httptest.ResponseRecorder

	// Try client context first, then login context
	if clientCtx != nil && clientCtx.GetStatusCode() != 0 {
		actualCode = clientCtx.GetStatusCode()
		response = clientCtx.GetResponse()
	} else if loginCtx != nil && loginCtx.GetStatusCode() != 0 {
		actualCode = loginCtx.GetStatusCode()
		response = loginCtx.GetResponse()
	} else {
		return fmt.Errorf("no active context with response found")
	}

	if actualCode != statusCode {
		responseStr := ""
		if response != nil {
			responseStr = response.Body.String()
		}
		return fmt.Errorf("expected status code %d, but got %d. Response: %s", statusCode, actualCode, responseStr)
	}
	return nil
}

// theResponseShouldContainAToken validates that the response contains a token
func theResponseShouldContainAToken() error {
	if loginCtx == nil {
		return fmt.Errorf("login context not initialized")
	}

	responseBody := loginCtx.GetResponseBody()
	if responseBody == nil {
		return fmt.Errorf("response body is nil")
	}

	token, ok := responseBody["Token"]
	if !ok || token == nil {
		return fmt.Errorf("response does not contain a token. Response: %v", responseBody)
	}

	tokenStr, ok := token.(string)
	if !ok || tokenStr == "" {
		return fmt.Errorf("token is not a valid string. Response: %v", responseBody)
	}

	return nil
}

// theResponseShouldContainUsername validates that the response contains a username
func theResponseShouldContainUsername() error {
	if loginCtx == nil {
		return fmt.Errorf("login context not initialized")
	}

	responseBody := loginCtx.GetResponseBody()
	if responseBody == nil {
		return fmt.Errorf("response body is nil")
	}

	username, ok := responseBody["Username"]
	if !ok || username == nil {
		return fmt.Errorf("response does not contain username. Response: %v", responseBody)
	}

	return nil
}

// theResponseShouldContainEmail validates that the response contains the expected email
func theResponseShouldContainEmail(email string) error {
	if loginCtx == nil {
		return fmt.Errorf("login context not initialized")
	}

	responseBody := loginCtx.GetResponseBody()
	if responseBody == nil {
		return fmt.Errorf("response body is nil")
	}

	responseEmail, ok := responseBody["Email"]
	if !ok {
		return fmt.Errorf("response does not contain email. Response: %v", responseBody)
	}

	emailStr, ok := responseEmail.(string)
	if !ok || emailStr != email {
		return fmt.Errorf("expected email %s, but got %v. Response: %v", email, responseEmail, responseBody)
	}

	return nil
}

// theResponseShouldContain validates that the response contains the expected message
func theResponseShouldContain(message string) error {
	var response *httptest.ResponseRecorder

	// Try client context first, then login context
	if clientCtx != nil && clientCtx.GetResponse() != nil {
		response = clientCtx.GetResponse()
	} else if loginCtx != nil && loginCtx.GetResponse() != nil {
		response = loginCtx.GetResponse()
	} else {
		return fmt.Errorf("no active context with response found")
	}

	if response == nil {
		return fmt.Errorf("response is nil")
	}

	responseStr := response.Body.String()
	if !bytes.Contains([]byte(responseStr), []byte(message)) {
		return fmt.Errorf("expected response to contain '%s', but got: %s", message, responseStr)
	}

	return nil
}
