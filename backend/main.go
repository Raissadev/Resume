package main

import (
	j "api/jobs"
	"api/src/config"
	"api/src/config/logger"
	"api/src/routes"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var ds config.DataSource

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Start crons
	go j.InitJobs()

	logger.Init()
	log.Println("Starting services...")
}

func main() {
	fmt.Println("Hello World, 42")

	// Recover in case err
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered. Error:\n", r)
		}
	}()

	// Conn of database
	if err := ds.Conn(); err != nil {
		logger.Error("[ERROR START DATABASE]", err)
	}
	defer ds.Close()

	routes := routes.New()

	routes.Mail()
	routes.NewWeb()

	// routes.Server.Listen(os.Getenv("SERVER_PORT"))

	if err := routes.Server.ListenTLS(":443", "/.private/fullchain.pem", "/.private/privkey.pem"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
