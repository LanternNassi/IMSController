package main

import (
	"log"

	"github.com/LanternNassi/IMSController/internal/database"
	"github.com/LanternNassi/IMSController/internal/server"
)

func main() {
	//Setting up the database

	client, err := database.NewDatabaseClient()

	migration_err := client.Migrate()

	if err != nil || migration_err != nil {
		log.Fatalf("Database not loaded ...")

	}

	server := server.NewEchoServer(client)

	starting_error := server.Start()

	if starting_error != nil {
		log.Fatal(starting_error.Error())
	}

}
