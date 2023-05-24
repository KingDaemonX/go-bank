package interfaces

import (
	"net/http"

	"github.com/KingDaemonX/evolve-mod-ddd-sample/adapt/responses"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/transactions/application"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/transactions/domain/entity"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TransacationAppInterface struct {
	ta  application.TransactionApplication
	res *responses.Response
	val *validator.Validate
}

func TransacationInterface(ta *application.TransactionApplication) *TransacationAppInterface {
	return &TransacationAppInterface{
		ta: *ta,
	}
}

func (ti *TransacationAppInterface) CreateTransfer() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userID")

		var transfer entity.Transfer

		if err := c.BindJSON(&transfer); err != nil {
			ti.res.ErrorResp(err)
		}

		if err := ti.val.Struct(&transfer); err != nil {
			ti.res.ErrorResp(err)
		}

		resp, err := ti.ta.CreateTransfer(userId, transfer)
		if err != nil {
			ti.res.ErrorResp(err, http.StatusInternalServerError)
		}

		ti.res.JSONResp(http.StatusCreated, "transaction successful", resp)
	}
}
