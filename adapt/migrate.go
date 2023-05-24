package adapt

import (
	accountEntity "github.com/KingDaemonX/evolve-mod-ddd-sample/services/accounts/domain/entity"
	transactionEntity "github.com/KingDaemonX/evolve-mod-ddd-sample/services/transactions/domain/entity"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/users/domain/entity"
	"gorm.io/gorm"
)

func autoMigrate(conn *gorm.DB) error {
	return conn.AutoMigrate(&entity.User{}, &accountEntity.Account{}, &transactionEntity.Transactions{})
}
