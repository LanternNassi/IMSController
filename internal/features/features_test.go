package features

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/cucumber/godog"
)

func TestFeatures(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	featureDir := filepath.Dir(filename)

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{featureDir},
			TestingT: t,
			Concurrency: 1,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func TestMain(m *testing.M) {
	status := m.Run()
	os.Exit(status)
}
