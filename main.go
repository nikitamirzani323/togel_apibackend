package main

import (
	"log"

	"bitbucket.org/isbtotogroup/apibackend_go/db"
	"bitbucket.org/isbtotogroup/apibackend_go/routes"
)

func main() {
	db.Init()
	app := routes.Init()
	log.Fatal(app.Listen(":7072"))
}
