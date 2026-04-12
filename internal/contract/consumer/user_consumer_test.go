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

// gormModelMatcher returns matchers for the embedded gorm.Model fields.
func gormModelMatcher() matchers.MapMatcher {
	return matchers.MapMatcher{
		"ID":        matchers.Like(1),
		"CreatedAt": matchers.Like("2024-01-01T00:00:00Z"),
		"UpdatedAt": matchers.Like("2024-01-01T00:00:00Z"),
	}
}

func userBody() matchers.MapMatcher {
	m := gormModelMatcher()
	m["Username"] = matchers.Like("johndoe")
	m["Email"] = matchers.Like("john@example.com")
	m["Verified"] = matchers.Like(false)
	return m
}

func TestAddUser(t *testing.T) {
	pact := newV2Pact(t)

	reqBody := map[string]interface{}{
		"Username": "johndoe",
		"Password": "secret123",
		"Email":    "john@example.com",
	}

	err := pact.
		AddInteraction().
		Given("no user with email john@example.com exists").
		UponReceiving("a POST request to register a user").
		WithRequest("POST", "/Users", func(b *consumer.V2RequestBuilder) {
			b.Header("Content-Type", matchers.Like("application/json"))
			b.JSONBody(matchers.MapMatcher{
				"Username": matchers.Like("johndoe"),
				"Password": matchers.Like("secret123"),
				"Email":    matchers.Like("john@example.com"),
			})
		}).
		WillRespondWith(201, func(b *consumer.V2ResponseBuilder) {
			b.Header("Content-Type", matchers.Like("application/json; charset=UTF-8"))
			b.JSONBody(userBody())
		}).
		ExecuteTest(t, func(cfg consumer.MockServerConfig) error {
			body, _ := json.Marshal(reqBody)
			resp, err := http.Post(mockURL(cfg, "/Users"), "application/json", bytes.NewReader(body))
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			require.Equal(t, http.StatusCreated, resp.StatusCode)
			return nil
		})

	require.NoError(t, err)
}

func TestLogin(t *testing.T) {
	pact := newV2Pact(t)

	reqBody := map[string]interface{}{
		"Email":    "john@example.com",
		"Password": "secret123",
	}

	err := pact.
		AddInteraction().
		Given("a user with email john@example.com exists").
		UponReceiving("a POST request to login").
		WithRequest("POST", "/Users/login", func(b *consumer.V2RequestBuilder) {
			b.Header("Content-Type", matchers.Like("application/json"))
			b.JSONBody(matchers.MapMatcher{
				"Email":    matchers.Like("john@example.com"),
				"Password": matchers.Like("secret123"),
			})
		}).
		WillRespondWith(202, func(b *consumer.V2ResponseBuilder) {
			b.Header("Content-Type", matchers.Like("application/json; charset=UTF-8"))
			b.JSONBody(matchers.MapMatcher{
				"Username": matchers.Like("johndoe"),
				"Token":    matchers.Like("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.test"),
				"Email":    matchers.Like("john@example.com"),
			})
		}).
		ExecuteTest(t, func(cfg consumer.MockServerConfig) error {
			body, _ := json.Marshal(reqBody)
			resp, err := http.Post(mockURL(cfg, "/Users/login"), "application/json", bytes.NewReader(body))
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			require.Equal(t, http.StatusAccepted, resp.StatusCode)
			return nil
		})

	require.NoError(t, err)
}

func TestGetUsers(t *testing.T) {
	pact := newV2Pact(t)

	err := pact.
		AddInteraction().
		Given("users exist").
		UponReceiving("a GET request for all users").
		WithRequest("GET", "/Users").
		WillRespondWith(200, func(b *consumer.V2ResponseBuilder) {
			b.Header("Content-Type", matchers.Like("application/json; charset=UTF-8"))
			b.JSONBody(matchers.EachLike(userBody(), 1))
		}).
		ExecuteTest(t, func(cfg consumer.MockServerConfig) error {
			resp, err := http.Get(mockURL(cfg, "/Users"))
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			require.Equal(t, http.StatusOK, resp.StatusCode)
			return nil
		})

	require.NoError(t, err)
}

func TestGetUserById(t *testing.T) {
	pact := newV2Pact(t)

	err := pact.
		AddInteraction().
		Given("a user with ID 1 exists").
		UponReceiving("a GET request for user by ID").
		WithRequest("GET", "/Users/1").
		WillRespondWith(200, func(b *consumer.V2ResponseBuilder) {
			b.Header("Content-Type", matchers.Like("application/json; charset=UTF-8"))
			b.JSONBody(userBody())
		}).
		ExecuteTest(t, func(cfg consumer.MockServerConfig) error {
			resp, err := http.Get(mockURL(cfg, "/Users/1"))
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			require.Equal(t, http.StatusOK, resp.StatusCode)
			return nil
		})

	require.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	pact := newV2Pact(t)

	err := pact.
		AddInteraction().
		Given("a user with ID 1 exists").
		UponReceiving("a DELETE request to remove a user").
		WithRequest("DELETE", "/Users/1").
		WillRespondWith(202, func(b *consumer.V2ResponseBuilder) {
			b.Header("Content-Type", matchers.Like("application/json; charset=UTF-8"))
			b.JSONBody(matchers.Like("User deleted successfully"))
		}).
		ExecuteTest(t, func(cfg consumer.MockServerConfig) error {
			req, err := http.NewRequest(http.MethodDelete, mockURL(cfg, "/Users/1"), nil)
			if err != nil {
				return err
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			require.Equal(t, http.StatusAccepted, resp.StatusCode)
			return nil
		})

	require.NoError(t, err)
}
