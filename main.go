package main

import (
	"log"
	"solid-spork/src/api"
	database "solid-spork/src/db"
)

func main() {
	db := database.InitDB()
	server := api.CreateServer(db)

	log.Fatal(server.Listen(":8080"))
}
