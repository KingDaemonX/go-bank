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
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error %s", err.Error())
	}
}

func main() {
	fmt.Println("Hello World From KingDaemonX \n This is the transfer / transaction route")

	service, err := adapt.ConnectDatabase()
	if err != nil {
		log.Fatalf("Error %s", err.Error())
	}

	route := gin.Default()

	route.Use(adapt.CorsMiddleWare())

	router.TransactionRouter(service, route)

	port := os.Getenv("TRANSACTION_PORT")
	if port == "" {
		port = "5000"
	}

	log.Printf("Starting Server At Localhost:%v", port)
	go log.Fatal(route.Run(":" + port))
	log.Printf("Running Server At Localhost:%v", port)
}
