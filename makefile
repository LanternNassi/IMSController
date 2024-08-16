# Variables
DOCKER_COMPOSE_FILE := docker-compose.yml
INFO := @echo 



# Build the Docker containers
build:
	@ ${INFO} "Building required docker images"
	@ docker-compose -f $(DOCKER_COMPOSE_FILE) build db go-test
	@ ${INFO} "Docker image built successfully"
	@ echo " "

# Run the Docker containers
run:
	@ ${INFO} "Running the Docker containers"
	@ docker-compose -f $(DOCKER_COMPOSE_FILE) up -d db go-test
	@ ${INFO} "Docker containers running successfully"
	@ echo " "

# Run the tests
test:run
	@ ${INFO} "Running tests"
	@ docker-compose -f $(DOCKER_COMPOSE_FILE) exec imscontroller-go-test-1 go test -v  
	@ ${INFO} "Tests completed successfully"
	@ echo " "

# Clean up Docker containers
clean:
	@ ${INFO} "Cleaning up Docker containers"
	@ docker-compose -f $(DOCKER_COMPOSE_FILE) down -v
	@ ${INFO} "Docker containers cleaned up successfully"

# Run clean, build, and test in sequence
rebuild-test: clean all
