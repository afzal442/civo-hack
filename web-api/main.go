package main

import (
	"REST-API/routes"
)

func main() {

	// Initialize the routes for you
	r := routes.InitializeRoutes()

	// Start serving the application
	r.Run()
}
