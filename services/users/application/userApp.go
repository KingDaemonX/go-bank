package application

import (
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/users/domain/entity"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/users/domain/repository"
)

type UserApplication struct {
	user repository.UserRepository
}

var _ UserAppInterface = &UserApplication{}

type UserAppInterface interface {
	CreateUser(*entity.User) (any, error)
	ReadUserByID(string) (any, error)
	UpdateUser(string, *entity.User) (any, error)
	LoginUser(*entity.UserLogin) (any, error)
}

func (ua *UserApplication) CreateUser(user *entity.User) (any, error) {
	return ua.user.CreateUser(user)
}

func (ua *UserApplication) LoginUser(user *entity.UserLogin) (any, error) {
	return ua.user.LoginUser(user)
}

func (ua *UserApplication) ReadUserByID(userId string) (any, error) {
	return ua.user.ReadUserByID(userId)
}

func (ua *UserApplication) UpdateUser(userId string, user *entity.User) (any, error) {
	return ua.user.UpdateUser(userId, user)
}
