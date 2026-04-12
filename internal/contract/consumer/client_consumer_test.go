package consumer_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/pact-foundation/pact-go/v2/consumer"
	"github.com/pact-foundation/pact-go/v2/matchers"
	"github.com/stretchr/testify/require"
)

// clientBody is a minimal matcher for a single client object in responses.
func clientBody() matchers.MapMatcher {
	return matchers.MapMatcher{
		"ClientID":     matchers.Like("abc123"),
		"FirstName":    matchers.Like("John"),
		"LastName":     matchers.Like("Doe"),
		"Email":        matchers.Like("john@example.com"),
		"Phone":        matchers.Like("+256700000000"),
		"Address":      matchers.Like("123 Main St"),
		"BusinessName": matchers.Like("ACME Corp"),
		"Status":       matchers.Like("connected"),
		"ValidTill":    matchers.Like("2025-01-01T00:00:00Z"),
	}
}

func TestGetClients(t *testing.T) {
	pact := newV2Pact(t)

	err := pact.
		AddInteraction().
		Given("clients exist").
		UponReceiving("a GET request for all clients").
		WithRequest("GET", "/clients").
		WillRespondWith(200, func(b *consumer.V2ResponseBuilder) {
			b.Header("Content-Type", matchers.Like("application/json; charset=UTF-8"))
			b.JSONBody(matchers.EachLike(clientBody(), 1))
		}).
		ExecuteTest(t, func(cfg consumer.MockServerConfig) error {
			resp, err := http.Get(mockURL(cfg, "/clients"))
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			require.Equal(t, http.StatusOK, resp.StatusCode)
			return nil
		})

	require.NoError(t, err)
}

func TestAddClient(t *testing.T) {
	pact := newV2Pact(t)

	reqBody := map[string]interface{}{
		"FirstName":    "John",
		"LastName":     "Doe",
		"Email":        "john@example.com",
		"Phone":        "+256700000000",
		"Address":      "123 Main St",
		"BusinessName": "ACME Corp",
		"Status":       "connected",
	}

	err := pact.
		AddInteraction().
		Given("a new client can be registered").
		UponReceiving("a POST request to create a client").
		WithRequest("POST", "/clients", func(b *consumer.V2RequestBuilder) {
			b.Header("Content-Type", matchers.Like("application/json"))
			b.JSONBody(matchers.MapMatcher{
				"FirstName":    matchers.Like("John"),
				"LastName":     matchers.Like("Doe"),
				"Email":        matchers.Like("john@example.com"),
				"Phone":        matchers.Like("+256700000000"),
				"Address":      matchers.Like("123 Main St"),
				"BusinessName": matchers.Like("ACME Corp"),
				"Status":       matchers.Like("connected"),
			})
		}).
		WillRespondWith(201, func(b *consumer.V2ResponseBuilder) {
			b.Header("Content-Type", matchers.Like("application/json; charset=UTF-8"))
			b.JSONBody(clientBody())
		}).
		ExecuteTest(t, func(cfg consumer.MockServerConfig) error {
			body, _ := json.Marshal(reqBody)
			resp, err := http.Post(mockURL(cfg, "/clients"), "application/json", bytes.NewReader(body))
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			require.Equal(t, http.StatusCreated, resp.StatusCode)
			return nil
		})

	require.NoError(t, err)
}

func TestGetClientById(t *testing.T) {
	pact := newV2Pact(t)

	err := pact.
		AddInteraction().
		Given("a client with ID test-client-id exists").
		UponReceiving("a GET request for client by ID").
		WithRequest("GET", "/clients/test-client-id").
		WillRespondWith(200, func(b *consumer.V2ResponseBuilder) {
			b.Header("Content-Type", matchers.Like("application/json; charset=UTF-8"))
			b.JSONBody(clientBody())
		}).
		ExecuteTest(t, func(cfg consumer.MockServerConfig) error {
			resp, err := http.Get(mockURL(cfg, "/clients/test-client-id"))
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			require.Equal(t, http.StatusOK, resp.StatusCode)
			return nil
		})

	require.NoError(t, err)
}

func TestUpdateClient(t *testing.T) {
	pact := newV2Pact(t)

	updateBody := map[string]interface{}{
		"FirstName":    "Jane",
		"LastName":     "Doe",
		"Email":        "jane@example.com",
		"Phone":        "+256700000001",
		"Address":      "456 New St",
		"BusinessName": "New Corp",
		"Status":       "connected",
		"ValidTill":    "2026-01-01T00:00:00Z",
	}

	err := pact.
		AddInteraction().
		Given("a client with ID test-client-id exists").
		UponReceiving("a PUT request to update a client").
		WithRequest("PUT", "/clients/test-client-id", func(b *consumer.V2RequestBuilder) {
			b.Header("Content-Type", matchers.Like("application/json"))
			b.JSONBody(matchers.MapMatcher{
				"FirstName":    matchers.Like("Jane"),
				"LastName":     matchers.Like("Doe"),
				"Email":        matchers.Like("jane@example.com"),
				"Phone":        matchers.Like("+256700000001"),
				"Address":      matchers.Like("456 New St"),
				"BusinessName": matchers.Like("New Corp"),
				"Status":       matchers.Like("connected"),
				"ValidTill":    matchers.Like("2026-01-01T00:00:00Z"),
			})
		}).
		WillRespondWith(200, func(b *consumer.V2ResponseBuilder) {
			b.Header("Content-Type", matchers.Like("application/json; charset=UTF-8"))
			b.JSONBody(clientBody())
		}).
		ExecuteTest(t, func(cfg consumer.MockServerConfig) error {
			body, _ := json.Marshal(updateBody)
			req, err := http.NewRequest(http.MethodPut, mockURL(cfg, "/clients/test-client-id"), bytes.NewReader(body))
			if err != nil {
				return err
			}
			req.Header.Set("Content-Type", "application/json")
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			require.Equal(t, http.StatusOK, resp.StatusCode)
			return nil
		})

	require.NoError(t, err)
}
