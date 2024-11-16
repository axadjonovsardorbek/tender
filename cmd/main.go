package main

import (
	"github.com/axadjonovsardorbek/tender/app"
	"github.com/axadjonovsardorbek/tender/config"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize the app
	app := &app.App{}
	app.Initialize(&cfg)
	defer app.Close()

	// Run the server
	app.Run(cfg.ServerPort)
}
