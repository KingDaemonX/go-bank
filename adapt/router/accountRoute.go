package router

import (
	"github.com/KingDaemonX/evolve-mod-ddd-sample/adapt"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/accounts/application"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/accounts/interfaces"
	"github.com/gin-gonic/gin"
)

func AccountRoute(service *adapt.ServiceInfra, router *gin.Engine) {
	account := interfaces.AccountInterface((*application.AccountApplication)(&service.AccountInfra))

	accounts := router.Group("/api/accounts")
	accounts.POST("/create", account.CreateAccount())
}
