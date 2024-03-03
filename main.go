package main

import (
	"assignment1/restfull_Server"
)

// main is the entry point of the application.
// It calls StartApplication from the restfull_Server package to initialize and start the server.
// This includes setting up all the routes, handlers, and any necessary initial configurations required for the server to run.
// The server will continue to run and serve requests until it's manually stopped or encounters a fatal error.
func main() {
	restfull_Server.StartApplication() // Initialize and start the server
}
