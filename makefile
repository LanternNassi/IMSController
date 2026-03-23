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
