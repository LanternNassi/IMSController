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

func installationBody() matchers.MapMatcher {
	m := gormModelMatcher()
	m["ClientID"] = matchers.Like("abc123")
	m["Installation_type"] = matchers.Like("server")
	m["Computer_name"] = matchers.Like("DESKTOP-001")
	m["IMS_version"] = matchers.Like("1.0.0")
	m["Operating_system"] = matchers.Like("Windows 11")
	m["RAM"] = matchers.Like("16GB")
	m["Processor"] = matchers.Like("Intel Core i7")
	m["Active"] = matchers.Like("true")
	return m
}

func TestGetInstallations(t *testing.T) {
	pact := newV2Pact(t)

	err := pact.
		AddInteraction().
		Given("installations exist").
		UponReceiving("a GET request for all installations").
		WithRequest("GET", "/Installations").
		WillRespondWith(200, func(b *consumer.V2ResponseBuilder) {
			b.Header("Content-Type", matchers.Like("application/json; charset=UTF-8"))
			b.JSONBody(matchers.EachLike(installationBody(), 1))
		}).
		ExecuteTest(t, func(cfg consumer.MockServerConfig) error {
			resp, err := http.Get(mockURL(cfg, "/Installations"))
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			require.Equal(t, http.StatusOK, resp.StatusCode)
			return nil
		})

	require.NoError(t, err)
}

func TestAddInstallation(t *testing.T) {
	pact := newV2Pact(t)

	reqBody := map[string]interface{}{
		"ClientID":          "abc123",
		"Installation_type": "server",
		"Computer_name":     "DESKTOP-001",
		"IMS_version":       "1.0.0",
		"Operating_system":  "Windows 11",
		"RAM":               "16GB",
		"Processor":         "Intel Core i7",
		"Active":            "true",
	}

	err := pact.
		AddInteraction().
		Given("a client with ID abc123 exists").
		UponReceiving("a POST request to create an installation").
		WithRequest("POST", "/Installations", func(b *consumer.V2RequestBuilder) {
			b.Header("Content-Type", matchers.Like("application/json"))
			b.JSONBody(matchers.MapMatcher{
				"ClientID":          matchers.Like("abc123"),
				"Installation_type": matchers.Like("server"),
				"Computer_name":     matchers.Like("DESKTOP-001"),
				"IMS_version":       matchers.Like("1.0.0"),
				"Operating_system":  matchers.Like("Windows 11"),
				"RAM":               matchers.Like("16GB"),
				"Processor":         matchers.Like("Intel Core i7"),
				"Active":            matchers.Like("true"),
			})
		}).
		WillRespondWith(201, func(b *consumer.V2ResponseBuilder) {
			b.Header("Content-Type", matchers.Like("application/json; charset=UTF-8"))
			b.JSONBody(installationBody())
		}).
		ExecuteTest(t, func(cfg consumer.MockServerConfig) error {
			body, _ := json.Marshal(reqBody)
			resp, err := http.Post(mockURL(cfg, "/Installations"), "application/json", bytes.NewReader(body))
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			require.Equal(t, http.StatusCreated, resp.StatusCode)
			return nil
		})

	require.NoError(t, err)
}

func TestGetInstallationById(t *testing.T) {
	pact := newV2Pact(t)

	err := pact.
		AddInteraction().
		Given("an installation with ID 1 exists").
		UponReceiving("a GET request for installation by ID").
		WithRequest("GET", "/Installations/1").
		WillRespondWith(200, func(b *consumer.V2ResponseBuilder) {
			b.Header("Content-Type", matchers.Like("application/json; charset=UTF-8"))
			b.JSONBody(installationBody())
		}).
		ExecuteTest(t, func(cfg consumer.MockServerConfig) error {
			resp, err := http.Get(mockURL(cfg, "/Installations/1"))
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			require.Equal(t, http.StatusOK, resp.StatusCode)
			return nil
		})

	require.NoError(t, err)
}
