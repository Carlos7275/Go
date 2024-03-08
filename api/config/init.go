package config

import (
	v1 "api/app/controllers/v1"
	"api/app/repositories"
	"api/app/services"
)

type Initialization struct {
	AuthSvc  services.AuthService
	AuthCtrl v1.AuthController
	UserRepo repositories.UserRepository
}

func NewInitialization(
	authService services.AuthService,
	authCtrl v1.AuthController,
	userRepo repositories.UserRepository,
) *Initialization {
	return &Initialization{
		AuthSvc:  authService,
		AuthCtrl: authCtrl,
		UserRepo: userRepo,
	}
}
