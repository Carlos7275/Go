//go:build wireinject
// +build wireinject

// go:build wireinject
package config

import (
	v1 "api/app/controllers/v1"
	"api/app/repositories"
	"api/app/services"

	"github.com/google/wire"
)

var db = wire.NewSet(DBConnection)

var authCtrlSet = wire.NewSet(
	v1.AuthControllerInit,
	wire.Bind(new(v1.AuthController), new(*v1.AuthControllerImpl)),
)

var userCtrlSet = wire.NewSet(
	v1.UserControllerInit,
	wire.Bind(new(v1.UserController), new(*v1.UserControllerImpl)),
)

var authServiceSet = wire.NewSet(
	services.NewAuthServiceImpl,
	wire.Bind(new(services.AuthService), new(*services.AuthServiceImpl)),
)

var userServiceSet = wire.NewSet(
	services.NewUserServiceImpl,
	wire.Bind(new(services.UserService), new(*services.UserServiceImpl)),
)

var userRepoSet = wire.NewSet(repositories.UserRepositoryInit,
	wire.Bind(new(repositories.UserRepository), new(*repositories.UserRepositoryImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, authCtrlSet, authServiceSet, userServiceSet, userRepoSet, userCtrlSet)
	return nil
}
