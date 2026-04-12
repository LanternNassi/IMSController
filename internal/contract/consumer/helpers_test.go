package consumer_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/pact-foundation/pact-go/v2/consumer"
	"github.com/stretchr/testify/require"
)

const (
	consumerName = "IMSFrontend"
	providerName = "IMSController"

	// testAuthHeader is a placeholder header value recorded in the pact file.
	// The provider test uses a RequestFilter to replace it with a valid JWT.
	testAuthHeader = "Bearer pact-test-token"
)

func pactDir() string {
	abs, _ := filepath.Abs(filepath.Join("..", "..", "..", "pacts"))
	return abs
}

func newV2Pact(t *testing.T) *consumer.V2HTTPMockProvider {
	t.Helper()
	pact, err := consumer.NewV2Pact(consumer.MockHTTPProviderConfig{
		Consumer: consumerName,
		Provider: providerName,
		PactDir:  pactDir(),
	})
	require.NoError(t, err)
	return pact
}

func mockURL(cfg consumer.MockServerConfig, path string) string {
	return fmt.Sprintf("http://%s:%d%s", cfg.Host, cfg.Port, path)
}
