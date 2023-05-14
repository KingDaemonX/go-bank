package auth

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			if err == http.ErrNoCookie {
				//come fix this later
				// responses.ErrorResp(c, err, http.StatusUnauthorized)
				return
			}
			// come fix this later
			// responses.ErrorResp(c, err, http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SIGNING_KEY")), nil
		})
		if err != nil {
			log.Panic(err)
			return
		}

		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			if claims.VerifyExpiresAt(time.Now(), true) {
				log.Printf("%v token has expired !!", claims.email)
				return
			}

			c.Next()
		}
	}
}
