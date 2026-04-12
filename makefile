# Variables
DOCKER_COMPOSE_FILE := docker-compose.yml
INFO := @echo 

# Build the Docker containers
buildTest:
	@ ${INFO} "Building required docker images for Testing"
	@ docker compose -f $(DOCKER_COMPOSE_FILE) build go_test test_db
	@ ${INFO} "Docker images built successfully"
	@ echo " "

buildAPI:
	@ ${INFO} "Building required docker images for the API"
	@ docker compose -f $(DOCKER_COMPOSE_FILE) build --no-cache db go_api
	@ ${INFO} "Docker images built successfully"
	@ echo " "


# Run the Docker containers
run:
	@ ${INFO} "Running the Docker containers"
	@ docker compose -f $(DOCKER_COMPOSE_FILE) up db go_api
	@ ${INFO} "Docker containers running successfully"
	@ echo " "

# Run the tests
test:
	@ ${INFO} "Running tests"
	@ docker compose -f $(DOCKER_COMPOSE_FILE) run --rm go_test go test -v -coverprofile=coverage.txt
	@ ${INFO} "Tests completed successfully"
	@ echo " "

# Run godog BDD tests
test-godog:
	@ ${INFO} "Running godog BDD tests"
	@ docker compose -f $(DOCKER_COMPOSE_FILE) run --rm go_test go test -v -count=1 ./internal/features
	@ ${INFO} "Godog tests completed successfully"
	@ echo " "

# Run godog BDD tests with coverage
test-godog-coverage:
	@ ${INFO} "Running godog BDD tests with coverage"
	@ docker compose -f $(DOCKER_COMPOSE_FILE) run --rm go_test go test -v -count=1 -coverprofile=coverage-godog.txt -coverpkg=./... ./internal/features
	@ ${INFO} "Godog tests with coverage completed successfully"
	@ echo " "

# Run godog tests locally (without Docker)
test-godog-local:
	@ ${INFO} "Running godog BDD tests locally"
	@ go test -v ./internal/features/...
	@ ${INFO} "Godog tests completed successfully"
	@ echo " "

# ── Pact contract testing ────────────────────────────────────────────────────
#
# pact-go v2 uses CGO to call the Rust Pact FFI library.
# The Docker targets below use the golang:latest builder image (which has gcc)
# and install the FFI automatically.
#
# To run locally on Linux/macOS, install the FFI first:
#   go run github.com/pact-foundation/pact-go/v2 install --libDir /usr/local/lib
# On Windows, use WSL2 or the Docker targets below.
# ─────────────────────────────────────────────────────────────────────────────

# Run Pact consumer tests inside Docker (generates pact files in ./pacts/)
test-pact-consumer:
	@ ${INFO} "Running Pact consumer tests (Docker)"
	@ docker compose -f $(DOCKER_COMPOSE_FILE) run --rm go_test sh -c \
		"go run github.com/pact-foundation/pact-go/v2 install --libDir /usr/local/lib && \
		 CGO_ENABLED=1 go test -v -count=1 ./internal/contract/consumer/..."
	@ ${INFO} "Consumer tests done — pact files written to ./pacts/"
	@ echo " "

# Run Pact provider verification inside Docker (requires pact files + test DB)
test-pact-provider:
	@ ${INFO} "Running Pact provider verification (Docker)"
	@ docker compose -f $(DOCKER_COMPOSE_FILE) run --rm go_test sh -c \
		"go run github.com/pact-foundation/pact-go/v2 install --libDir /usr/local/lib && \
		 CGO_ENABLED=1 go test -v -count=1 ./internal/contract/provider/..."
	@ ${INFO} "Provider verification completed"
	@ echo " "

# Full Pact workflow: generate pacts then verify provider
test-pact: test-pact-consumer test-pact-provider
	@ ${INFO} "Pact contract testing completed"
	@ echo " "


# Run godog tests locally with coverage
test-godog-local-coverage:
	@ ${INFO} "Running godog BDD tests locally with coverage"
	@ go test -v -coverprofile=coverage-godog.txt -coverpkg=./... ./internal/features/...
	@ ${INFO} "Coverage report generated: coverage-godog.txt"
	@ ${INFO} "View coverage with: go tool cover -html=coverage-godog.txt"
	@ echo " "

# Generate HTML coverage report
coverage-html:
	@ ${INFO} "Generating HTML coverage report"
	@ if [ -f coverage.txt ]; then \
		go tool cover -html=coverage.txt -o coverage.html; \
		echo "Coverage report generated: coverage.html"; \
	elif [ -f coverage-godog.txt ]; then \
		go tool cover -html=coverage-godog.txt -o coverage.html; \
		echo "Coverage report generated: coverage.html"; \
	else \
		echo "No coverage file found. Run 'make test' or 'make test-godog-coverage' first"; \
	fi
	@ echo " "

# Clean up Docker containers
clean:
	@ ${INFO} "Cleaning up Docker containers"
	@ docker compose -f $(DOCKER_COMPOSE_FILE) down -v
	@ ${INFO} "Docker containers cleaned up successfully"

# Run clean, build, and test in sequence
rebuild-test: clean build test
