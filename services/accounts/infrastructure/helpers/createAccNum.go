package helpers

import (
	// "crypto/rand"
	// "math/big"

	"crypto/rand"
	"log"
	"math/big"

	"github.com/KingDaemonX/evolve-mod-ddd-sample/adapt"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/accounts/domain/entity"
)

type AccountHelper struct {
	account *entity.Account
}

func (ah *AccountHelper) GenerateAccNumber() string {
	for {
		accNum, err := rand.Int(rand.Reader, big.NewInt(10000000000))
		if err != nil {
			log.Fatalf("Error %s", err.Error())
		}

		if len(accNum.String()) != 10 {
			ah.GenerateAccNumber()
			continue
		}

		valid := ah.CheckAccountNumber(accNum.String())

		if !valid {
			ah.GenerateAccNumber()
			continue
		}
		return accNum.String()
	}
}

func (ah *AccountHelper) CheckAccountNumber(accNum string) bool {
	isValid := false
	if err := adapt.Conn.Debug().Find(ah.account, "account_number ?=", accNum); err != nil {
		log.Println(err.Error)
		return isValid
	}

	if ah.account.UID > 0 {
		return isValid
	}

	isValid = true
	return isValid
}
