package main

import (
	"fmt"
	"log"
	"os"

	"github.com/KingDaemonX/evolve-mod-ddd-sample/adapt"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/adapt/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// load .env variable here
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error %s", err.Error())
	}
}

func main() {
	fmt.Println("Hello World from KingDaemonX\n this is the account microservice route")

	service, err := adapt.ConnectDatabase()
	if err != nil {
		log.Fatalf("Error %s", err.Error())
	}

	route := gin.Default()

	router.AccountRoute(service, route)

	port := os.Getenv("ACCOUNT_PORT")
	if port == "" {
		port = "3001"
	}

	log.Printf("Starting Server At Localhost:%v", port)
	go log.Fatal(route.Run(":" + port))
	log.Printf("Running Server At Localhost:%v", port)
}
