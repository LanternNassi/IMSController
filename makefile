# Variables
DOCKER_COMPOSE_FILE := docker-compose.yml
INFO := @echo 

# Build the Docker containers
buildTest:
	@ ${INFO} "Building required docker images for Testing"
	@ docker-compose -f $(DOCKER_COMPOSE_FILE) build go_test test_db
	@ ${INFO} "Docker images built successfully"
	@ echo " "

buildAPI:
	@ ${INFO} "Building required docker images for the API"
	@ docker-compose -f $(DOCKER_COMPOSE_FILE) build --no-cache db go_api
	@ ${INFO} "Docker images built successfully"
	@ echo " "


# Run the Docker containers
run:
	@ ${INFO} "Running the Docker containers"
	@ docker-compose -f $(DOCKER_COMPOSE_FILE) up db go_api
	@ ${INFO} "Docker containers running successfully"
	@ echo " "

# Run the tests
test:
	@ ${INFO} "Running tests"
	@ docker-compose -f $(DOCKER_COMPOSE_FILE) run --rm go_test go test -v -coverprofile=coverage.txt
	@ ${INFO} "Tests completed successfully"
	@ echo " "

# Clean up Docker containers
clean:
	@ ${INFO} "Cleaning up Docker containers"
	@ docker-compose -f $(DOCKER_COMPOSE_FILE) down -v
	@ ${INFO} "Docker containers cleaned up successfully"

# Run clean, build, and test in sequence
rebuild-test: clean build test
