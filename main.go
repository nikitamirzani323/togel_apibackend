package main

import (
	"log"

	"github.com/nikitamirzani323/togel_apibackend/db"
	"github.com/nikitamirzani323/togel_apibackend/routes"
)

func main() {
	db.Init()
	app := routes.Init()
	log.Fatal(app.Listen(":7072"))
}
