package main

import (
	"REST-API/routes"
)

func main() {

	// Initialize the routes for you gh
	r := routes.InitializeRoutes()

	// Start serving the application
	r.Run()
}
