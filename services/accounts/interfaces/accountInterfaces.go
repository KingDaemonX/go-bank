package interfaces

import (
	"net/http"

	"github.com/KingDaemonX/evolve-mod-ddd-sample/adapt/responses"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/accounts/application"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/accounts/domain/entity"
	"github.com/gin-gonic/gin"
)

type AccAppInterface struct {
	aa  application.AccountApplication
	res responses.Response
}

func AccountInterface(aa *application.AccountApplication) *AccAppInterface {
	return &AccAppInterface{aa: *aa}
}

func (ai *AccAppInterface) CreateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var accounts *entity.Account

		if err := c.BindJSON(&accounts); err != nil {
			ai.res.ErrorResp(err)
		}

		resp, err := ai.aa.CreateAccount(accounts)
		if err != nil {
			ai.res.ErrorResp(err)
		}

		ai.res.JSONResp(http.StatusCreated, "Account Created", resp)
	}
}

func (ai *AccAppInterface) ReadAccountByAccountNumber() gin.HandlerFunc {
	return func(c *gin.Context) {
		accountNumber := c.Query("account-number")

		accountInfo, err := ai.aa.ReadAccountByAccountNumber(accountNumber)
		if err != nil {
			ai.res.ErrorResp(err, http.StatusNotFound)
		}

		ai.res.JSONResp(http.StatusOK, "success", accountInfo)
	}
}
