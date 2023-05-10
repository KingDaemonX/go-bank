package infrastructure

import (
	"errors"
	"strings"

	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/users/domain/entity"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/users/domain/repository"
	"gorm.io/gorm"
)

type UserInfra struct {
	Conn *gorm.DB
}

func NewUserInfra(Conn *gorm.DB) *UserInfra {
	return &UserInfra{Conn: Conn}
}

var _ repository.UserRepository = &UserInfra{}

func (u *UserInfra) CreateUser(user *entity.User) (any, error) {
	if err := u.Conn.Debug().Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return nil, errors.New("user with this profile already exist")
		}
		return nil, err
	}
	return user, nil
}

func (u *UserInfra) ReadUserByID(userId string) (any, error) {
	var user entity.User
	if err := u.Conn.Debug().Find(&user, "id ?=", userId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user record not found")
		}
		return nil, err
	}
	// do a public profile response : Full name, and account number
	return user, nil
}

func (u *UserInfra) UpdateUser(userId string, user *entity.User) (any, error) {
	return "", nil
}
