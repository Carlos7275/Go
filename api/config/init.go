package config

import (
	v1 "api/app/controllers/v1"
	"api/app/repositories"
	"api/app/services"
)

type Initialization struct {
	AuthSvc  services.AuthService
	UserSvc  services.UserService
	AuthCtrl v1.AuthController
	UserCtrl v1.UserController
	UserRepo repositories.UserRepository
}

func NewInitialization(
	authService services.AuthService,
	authCtrl v1.AuthController,
	userCtrl v1.UserController,
	userRepo repositories.UserRepository,
	userService services.UserService,
) *Initialization {
	return &Initialization{
		AuthSvc:  authService,
		AuthCtrl: authCtrl,
		UserRepo: userRepo,
		UserCtrl: userCtrl,
		UserSvc:  userService,
	}
}
