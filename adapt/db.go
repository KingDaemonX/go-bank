package adapt

import (
	"log"
	"os"

	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/users/domain/repository"
	infrastructure "github.com/KingDaemonX/evolve-mod-ddd-sample/services/users/infrastructure/persistent"
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
	UserInfra repository.UserRepository
	db        *gorm.DB
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
		UserInfra: infrastructure.NewUserInfra(db),
		db:        db,
	}, nil
}
