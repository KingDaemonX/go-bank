package router

import (
	"github.com/KingDaemonX/evolve-mod-ddd-sample/adapt"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/transactions/application"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/transactions/interfaces"
	"github.com/gin-gonic/gin"
)

func TransactionRouter(service *adapt.ServiceInfra, router *gin.Engine) {
	transactions := interfaces.TransacationInterface((*application.TransactionApplication)(&service.TransactionInfra))

	transaction := router.Group("/api/transaction")
	transaction.POST("/transfer", transactions.CreateTransfer())
}
