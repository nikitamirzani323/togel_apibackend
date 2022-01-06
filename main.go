package main

import (
	"log"
	"os"

	"bitbucket.org/isbtotogroup/apibackend_go/db"
	"bitbucket.org/isbtotogroup/apibackend_go/routes"
)

func main() {
	db.Init()
	app := routes.Init()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT env is required")
	}
	log.Fatal(app.Listen(":" + port))
}
