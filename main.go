package main

import (
	"log"

	"github.com/LanternNassi/IMSController/internal/database"
	"github.com/LanternNassi/IMSController/internal/server"
)

func main() {
	//Setting up the database

	client, err := database.NewDatabaseClient()

	// mongo_client, mongo_err := database.NewMongoDatabaseClient()

	// //Connecting to mongo database
	// if mongo_err != nil {
	// 	log.Fatal(mongo_err.Error())
	// }

	// mongo_object, mongo_context, mongo_error := mongo_client.ConnectMongo("mongodb+srv://lanternnassi:3YK3OyDBinRGRlzB@cluster0.15lo0el.mongodb.net/?retryWrites=true&w=majority")

	// if mongo_error != nil {
	// 	log.Fatal(mongo_error.Error())
	// }

	// mongo_ping_error := mongo_client.PingMongo(mongo_object, mongo_context)

	// if mongo_ping_error != nil {
	// 	log.Fatal(mongo_ping_error)
	// }

	//Migrating the postgres database
	migration_err := client.Migrate()

	if err != nil || migration_err != nil {
		log.Fatalf("Database not loaded ...")

	}

	server := server.NewEchoServer(client)

	starting_error := server.Start()

	if starting_error != nil {
		log.Fatal(starting_error.Error())
	}

	// defer mongo_client.CloseMongo(mongo_object, mongo_context)

}
