package auth

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/users/domain/entity"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type MyCustomClaims struct {
	UID       uint
	firstName string
	lastName  string
	email     string
	jwt.RegisteredClaims
}

func GenerateCookie(user *entity.User) {
	var c *gin.Context

	// Create claims with multiple fields populated
	claims := MyCustomClaims{
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "BankOfDaemon",
			Subject:   "KingDaemonX",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(os.Getenv("SIGNING_KEY"))
	if err != nil {
		log.Println("Error : ", err.Error())
		return
	}

	c.SetSameSite(http.SameSiteDefaultMode)
	c.SetCookie("Authorization", tokenString, 12*int(time.Hour), "/api", "", false, true)
	log.Println("Cookie Generated successfully")
}
