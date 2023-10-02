package main

import (
	"helloion/cmd"
	"helloion/database"
	"helloion/utils"
	"log"
)

func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Initialize Database
	database.Init(&config)

	// Execute command
	cmd.Execute()
}
