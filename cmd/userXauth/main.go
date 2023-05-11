package main

import (
	"fmt"
	"log"
	"os"

	"github.com/KingDaemonX/evolve-mod-ddd-sample/adapt"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/adapt/router"
	"github.com/gin-gonic/gin"
)

func init() {
	// load .env variable here
}

func main() {
	fmt.Println("Hello World From KingDaemonX 👋 \nI am building a bank application")

	services, err := adapt.ConnectDatabase()
	if err != nil {
		log.Panic(err)
		return
	}

	// user := interfaces.NewUserInterface((*application.UserAppInterface)(&services.UserInfra))

	route := gin.Default()

	route.Use(adapt.CorsMiddleWare())

	router.UserRouter(route, services)

	userPort := os.Getenv("USER_PORT")
	if userPort == "" {
		userPort = "3000"
	}
	log.Printf("Starting Server At Localhost:%v", userPort)
	log.Fatal(route.Run(":" + userPort))
	log.Printf("Running Server At Localhost:%v", userPort)
}
