package interfaces

import (
	"log"
	"net/http"

	"github.com/KingDaemonX/evolve-mod-ddd-sample/adapt/responses"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/users/application"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/users/domain/entity"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userInterface struct {
	userApp application.UserAppInterface
}

func NewUserInterface(ua *application.UserAppInterface) *userInterface {
	return &userInterface{userApp: *ua}
}

var validate = validator.New()

func (ui *userInterface) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user entity.User

		if err := c.BindJSON(&user); err != nil {
			responses.ErrorResp(c, err)
			log.Println("Error : ", err.Error())
			return
		}

		if err := validate.Struct(&user); err != nil {
			responses.ErrorResp(c, err)
			log.Println("Error : ", err.Error())
			return
		}

		resp, err := ui.userApp.CreateUser(&user)
		if err != nil {
			log.Println("Error : ", err.Error())
			return
		}

		responses.JSONResp(c, http.StatusCreated, "success", resp)
	}
}

func (ui *userInterface) GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")

		resp, err := ui.userApp.ReadUserByID(userID)
		if err != nil {
			log.Fatalf("Error %s", err.Error())
		}

		responses.JSONResp(c, http.StatusOK, "success", resp)
	}
}
