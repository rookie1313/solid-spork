package main

import (
	"log"
	database "solid-spork/src/db"
	"solid-spork/src/server"
)

func main() {
	db := database.InitDB()
	app := server.CreateApp(db)

	log.Fatal(app.Listen(":8080"))
}
