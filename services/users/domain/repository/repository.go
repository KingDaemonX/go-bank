package repository

import "github.com/KingDaemonX/evolve-mod-ddd-sample/services/users/domain/entity"

type UserRepository interface {
	CreateUser(*entity.User) (any, error)
	ReadUserByID(string) (any, error)
	UpdateUser(string, *entity.User) (any, error)
}
