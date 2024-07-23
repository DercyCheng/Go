//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"gt-go/webook/wire/repository"
	"gt-go/webook/wire/repository/dao"
)

func InitUserRepository() *repository.UserRepository {
	wire.Build(repository.NewUserRepository, InitDB, dao.NewUserDAO)
	return &repository.UserRepository{}
}
