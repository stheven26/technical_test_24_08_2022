package main

import (
	"github.com/stheven26/db"

	"github.com/stheven26/routes"
)

func main() {
	db.SetupDB()
	routes.StartApplication()
}
