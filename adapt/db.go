package adapt

import (
	"log"
	"os"

	accountRepo "github.com/KingDaemonX/evolve-mod-ddd-sample/services/accounts/domain/repository"
	accountPersistent "github.com/KingDaemonX/evolve-mod-ddd-sample/services/accounts/infrastructure/persistent"
	transactionRepo "github.com/KingDaemonX/evolve-mod-ddd-sample/services/transactions/domain/repository"
	transaactionPersistent "github.com/KingDaemonX/evolve-mod-ddd-sample/services/transactions/infrastructure/persistent"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/users/domain/repository"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/users/infrastructure/persistent"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

const (
	connectingDB   string = "🌀 Attempting To Connect Application Database"
	errDBConn      string = "🚨 Error Occur While Connecting To Database"
	successfulConn string = "😎 Database Connected SuccessFully"
	migrationErr   string = "❎ Error Occur While Migrating Database Schema"
)

type ServiceInfra struct {
	UserInfra        repository.UserRepository
	AccountInfra     accountRepo.AcccountRepository
	TransactionInfra transactionRepo.TransactionRepository
	db               *gorm.DB
}

func ConnectDatabase() (*ServiceInfra, error) {
	var err error
	log.Println(connectingDB)

	dsn := os.Getenv("DATABASE_CONN_STRING")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(errDBConn)
		return nil, err
	}

	Conn = db
	if err = autoMigrate(Conn); err != nil {
		log.Println(migrationErr)
		return nil, err
	}

	return &ServiceInfra{
		UserInfra:        persistent.NewUserInfra(db),
		AccountInfra:     accountPersistent.NewAccountInfra(db),
		TransactionInfra: transaactionPersistent.NewTrasactionInfra(db),
		db:               db,
	}, nil
}
