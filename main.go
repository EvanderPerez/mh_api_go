package main

import (
	"fmt"
	"mh_api_go/database"
	"mh_api_go/database/seed"
	"mh_api_go/routes"
	"os"
)

func main() {
	database.ConnectDatabase()

	// Check command-line arguments
	if len(os.Args) > 1 && os.Args[1] == "reset" {
		seed.SeedDatabase()
	} else {
		fmt.Println("Run with 'go run main.go reset' to reset and seed the database.")
	}

	r := routes.SetupRouter()
	r.Run(":3000")
}
