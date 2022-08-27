package main

import (
	"technical_test_24_08_2022/db"

	"technical_test_24_08_2022/routes"
)

func main() {
	db.SetupDB()
	routes.StartApplication()
}
