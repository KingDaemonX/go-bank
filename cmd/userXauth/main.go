package main

import (
	"fmt"
	"log"

	"github.com/KingDaemonX/evolve-mod-ddd-sample/adapt"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/users/application"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/users/interfaces"
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

	user := interfaces.NewUserInterface((*application.UserAppInterface)(&services.UserInfra))

	
}
