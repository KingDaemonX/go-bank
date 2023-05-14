package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

var c *gin.Context

func (res *Response) ErrorResp(err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	c.JSON(statusCode, Response{Status: statusCode, Message: "error", Data: err.Error()})
}

// resturn json response
func (res *Response) JSONResp(status int, message string, data any) {
	c.JSON(status, Response{Status: status, Message: message, Data: data})
}
