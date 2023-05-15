package adapt

import (
	"log"
	"os"

	accRepository "github.com/KingDaemonX/evolve-mod-ddd-sample/services/accounts/domain/repository"
	accInfrastructure "github.com/KingDaemonX/evolve-mod-ddd-sample/services/accounts/infrastructure/persistent"
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

type UserInfra struct {
	UserInfra    repository.UserRepository
	AccountInfra accRepository.AcccountRepository
	db           *gorm.DB
}

func ConnectDatabase() (*UserInfra, error) {
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

	return &UserInfra{
		UserInfra:    persistent.NewUserInfra(db),
		AccountInfra: accInfrastructure.NewAccountInfra(db),
		db:           db,
	}, nil
}
