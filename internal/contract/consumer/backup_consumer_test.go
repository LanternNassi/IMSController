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

// billSummary returns a minimal bill matcher for embedding inside backup responses.
func billSummary() matchers.MapMatcher {
	m := gormModelMatcher()
	m["ClientID"] = matchers.Like("abc123")
	m["BackupCount"] = matchers.Like(1)
	m["BackupSize"] = matchers.Like(1024)
	m["TotalCost"] = matchers.Like("0.0018273998877")
	m["Billed"] = matchers.Like(false)
	return m
}

func backupBody() matchers.MapMatcher {
	m := gormModelMatcher()
	m["ClientID"] = matchers.Like("abc123")
	m["Name"] = matchers.Like("daily-backup-2024")
	m["Backup"] = matchers.Like("s3://bucket/backup.tar.gz")
	m["Size"] = matchers.Like(1024)
	m["BillID"] = matchers.Like(1)
	m["Bill"] = matchers.Like(billSummary())
	return m
}

// withAuthHeader adds the test Authorization header to a request builder.
func withAuthHeader(b *consumer.V2RequestBuilder) {
	b.Header("Authorization", matchers.Like(testAuthHeader))
}

func TestGetBackups(t *testing.T) {
	pact := newV2Pact(t)

	err := pact.
		AddInteraction().
		Given("backups exist").
		UponReceiving("an authenticated GET request for all backups").
		WithRequest("GET", "/backups", func(b *consumer.V2RequestBuilder) {
			withAuthHeader(b)
		}).
		WillRespondWith(200, func(b *consumer.V2ResponseBuilder) {
			b.Header("Content-Type", matchers.Like("application/json; charset=UTF-8"))
			b.JSONBody(matchers.EachLike(backupBody(), 1))
		}).
		ExecuteTest(t, func(cfg consumer.MockServerConfig) error {
			req, err := http.NewRequest(http.MethodGet, mockURL(cfg, "/backups"), nil)
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

func TestAddBackup(t *testing.T) {
	pact := newV2Pact(t)

	reqBody := map[string]interface{}{
		"ClientID": "abc123",
		"Name":     "daily-backup-2024",
		"Backup":   "s3://bucket/backup.tar.gz",
		"Size":     1024,
	}

	err := pact.
		AddInteraction().
		Given("a client with ID abc123 exists").
		UponReceiving("an authenticated POST request to create a backup").
		WithRequest("POST", "/backups", func(b *consumer.V2RequestBuilder) {
			withAuthHeader(b)
			b.Header("Content-Type", matchers.Like("application/json"))
			b.JSONBody(matchers.MapMatcher{
				"ClientID": matchers.Like("abc123"),
				"Name":     matchers.Like("daily-backup-2024"),
				"Backup":   matchers.Like("s3://bucket/backup.tar.gz"),
				"Size":     matchers.Like(1024),
			})
		}).
		WillRespondWith(201, func(b *consumer.V2ResponseBuilder) {
			b.Header("Content-Type", matchers.Like("application/json; charset=UTF-8"))
			b.JSONBody(backupBody())
		}).
		ExecuteTest(t, func(cfg consumer.MockServerConfig) error {
			body, _ := json.Marshal(reqBody)
			req, err := http.NewRequest(http.MethodPost, mockURL(cfg, "/backups"), bytes.NewReader(body))
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

func TestGetBackupById(t *testing.T) {
	pact := newV2Pact(t)

	err := pact.
		AddInteraction().
		Given("a backup with ID 1 exists").
		UponReceiving("an authenticated GET request for backup by ID").
		WithRequest("GET", "/backups/1", func(b *consumer.V2RequestBuilder) {
			withAuthHeader(b)
		}).
		WillRespondWith(200, func(b *consumer.V2ResponseBuilder) {
			b.Header("Content-Type", matchers.Like("application/json; charset=UTF-8"))
			b.JSONBody(backupBody())
		}).
		ExecuteTest(t, func(cfg consumer.MockServerConfig) error {
			req, err := http.NewRequest(http.MethodGet, mockURL(cfg, "/backups/1"), nil)
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

func TestDeleteBackupById(t *testing.T) {
	pact := newV2Pact(t)

	err := pact.
		AddInteraction().
		Given("a backup with ID 1 exists").
		UponReceiving("an authenticated DELETE request to remove a backup").
		WithRequest("DELETE", "/backups/delete/1", func(b *consumer.V2RequestBuilder) {
			withAuthHeader(b)
		}).
		WillRespondWith(200, func(b *consumer.V2ResponseBuilder) {
			b.Header("Content-Type", matchers.Like("application/json; charset=UTF-8"))
			b.JSONBody(matchers.Like(true))
		}).
		ExecuteTest(t, func(cfg consumer.MockServerConfig) error {
			req, err := http.NewRequest(http.MethodDelete, mockURL(cfg, "/backups/delete/1"), nil)
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
