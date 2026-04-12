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

func billBody() matchers.MapMatcher {
	m := gormModelMatcher()
	m["ClientID"] = matchers.Like("abc123")
	m["BackupCount"] = matchers.Like(1)
	m["BackupSize"] = matchers.Like(1024)
	m["TotalCost"] = matchers.Like("0")
	m["Billed"] = matchers.Like(false)
	return m
}

func TestGetBills(t *testing.T) {
	pact := newV2Pact(t)

	err := pact.
		AddInteraction().
		Given("bills exist").
		UponReceiving("an authenticated GET request for all bills").
		WithRequest("GET", "/bills", func(b *consumer.V2RequestBuilder) {
			withAuthHeader(b)
		}).
		WillRespondWith(200, func(b *consumer.V2ResponseBuilder) {
			b.Header("Content-Type", matchers.Like("application/json; charset=UTF-8"))
			b.JSONBody(matchers.EachLike(billBody(), 1))
		}).
		ExecuteTest(t, func(cfg consumer.MockServerConfig) error {
			req, err := http.NewRequest(http.MethodGet, mockURL(cfg, "/bills"), nil)
			if err != nil {
				return err
			}
			req.Header.Set("Authorization", testAuthHeader)
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

func TestAddBill(t *testing.T) {
	pact := newV2Pact(t)

	reqBody := map[string]interface{}{
		"ClientID": "abc123",
	}

	err := pact.
		AddInteraction().
		Given("a client with ID abc123 exists").
		UponReceiving("an authenticated POST request to create a bill").
		WithRequest("POST", "/bills", func(b *consumer.V2RequestBuilder) {
			withAuthHeader(b)
			b.Header("Content-Type", matchers.Like("application/json"))
			b.JSONBody(matchers.MapMatcher{
				"ClientID": matchers.Like("abc123"),
			})
		}).
		WillRespondWith(201, func(b *consumer.V2ResponseBuilder) {
			b.Header("Content-Type", matchers.Like("application/json; charset=UTF-8"))
			b.JSONBody(billBody())
		}).
		ExecuteTest(t, func(cfg consumer.MockServerConfig) error {
			body, _ := json.Marshal(reqBody)
			req, err := http.NewRequest(http.MethodPost, mockURL(cfg, "/bills"), bytes.NewReader(body))
			if err != nil {
				return err
			}
			req.Header.Set("Authorization", testAuthHeader)
			req.Header.Set("Content-Type", "application/json")
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			require.Equal(t, http.StatusCreated, resp.StatusCode)
			return nil
		})

	require.NoError(t, err)
}

func TestGetBillById(t *testing.T) {
	pact := newV2Pact(t)

	err := pact.
		AddInteraction().
		Given("a bill with ID 1 exists").
		UponReceiving("an authenticated GET request for bill by ID").
		WithRequest("GET", "/bills/1", func(b *consumer.V2RequestBuilder) {
			withAuthHeader(b)
		}).
		WillRespondWith(200, func(b *consumer.V2ResponseBuilder) {
			b.Header("Content-Type", matchers.Like("application/json; charset=UTF-8"))
			b.JSONBody(billBody())
		}).
		ExecuteTest(t, func(cfg consumer.MockServerConfig) error {
			req, err := http.NewRequest(http.MethodGet, mockURL(cfg, "/bills/1"), nil)
			if err != nil {
				return err
			}
			req.Header.Set("Authorization", testAuthHeader)
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

func TestGetBillsByClientId(t *testing.T) {
	pact := newV2Pact(t)

	err := pact.
		AddInteraction().
		Given("bills exist for client abc123").
		UponReceiving("an authenticated GET request for bills by client ID").
		WithRequest("GET", "/bills/client/abc123", func(b *consumer.V2RequestBuilder) {
			withAuthHeader(b)
		}).
		WillRespondWith(200, func(b *consumer.V2ResponseBuilder) {
			b.Header("Content-Type", matchers.Like("application/json; charset=UTF-8"))
			b.JSONBody(matchers.EachLike(billBody(), 1))
		}).
		ExecuteTest(t, func(cfg consumer.MockServerConfig) error {
			req, err := http.NewRequest(http.MethodGet, mockURL(cfg, "/bills/client/abc123"), nil)
			if err != nil {
				return err
			}
			req.Header.Set("Authorization", testAuthHeader)
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
