package router

import (
	"github.com/KingDaemonX/evolve-mod-ddd-sample/adapt"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/users/application"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/users/interfaces"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine, services *adapt.ServiceInfra) {
	user := interfaces.NewUserInterface((*application.UserAppInterface)(&services.UserInfra))

	api := r.Group("/api")
	api.GET("/users/auth/login", user.LoginUser())

	api.POST("/users", user.CreateUser())
	api.GET("/users/:userId", user.GetUserByID())
}
