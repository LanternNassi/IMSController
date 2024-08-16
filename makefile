# Variables
DOCKER_COMPOSE_FILE := docker-compose.yml


# Build the Docker containers
build:
	@ ${INFO} "Building required docker images"
	@docker-compose -f $(DOCKER_COMPOSE_FILE) build db go-test
	@ ${INFO} "Docker image built successfully"
	@ echo " "

# Run the tests
test:
	@ ${INFO} "Running tests"
	@docker-compose -f $(DOCKER_COMPOSE_FILE) run --rm go-test -v 
	@ ${INFO} "Tests completed successfully"
	@ echo " "

# Clean up Docker containers
clean:
	@ ${INFO} "Cleaning up Docker containers"
	@docker-compose -f $(DOCKER_COMPOSE_FILE) down -v
	@ ${INFO} "Docker containers cleaned up successfully"

# Run clean, build, and test in sequence
rebuild-test: clean all
