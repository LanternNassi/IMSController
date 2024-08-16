# Variables
DOCKER_COMPOSE_FILE=docker-compose.yml

# Targets
.PHONY: all build test clean

# Default target
all: build test

# Build the Docker containers
build:
	@docker-compose -f $(DOCKER_COMPOSE_FILE) up --build -d

# Run the tests
test:
	@docker-compose -f $(DOCKER_COMPOSE_FILE) run --rm go-test -v 

# Clean up Docker containers
clean:
	@docker-compose -f $(DOCKER_COMPOSE_FILE) down -v

# Run clean, build, and test in sequence
rebuild-test: clean all
