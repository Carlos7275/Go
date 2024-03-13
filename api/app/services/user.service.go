package services

import (
	"api/app/models/dto"
	"api/app/models/requests"
	"api/app/repositories"
)

type UserService interface {
	AddUser(requests.UserRequest)
	ModifyUser(requests.User)
	FindUser(id int) (*dto.UsuariosDTO, error)
	GetUsers() ([]dto.UsuariosDTO, error)
	ChangeUserStatus(id int)
}

type UserServiceImpl struct {
	userRepository repositories.UserRepository
}

func (s UserServiceImpl) AddUser(requests.UserRequest) {

}

func (s UserServiceImpl) ModifyUser(requests.User) {

}
func (s UserServiceImpl) FindUser(id int) (*dto.UsuariosDTO, error) {
	user, err := s.userRepository.FindUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s UserServiceImpl) GetUsers() ([]dto.UsuariosDTO, error) {
	users, err := s.userRepository.FindAllUser()

	if err != nil {
		return nil, err
	}
	return users, nil
}
func (s UserServiceImpl) ChangeUserStatus(id int) {

}

func NewUserServiceImpl(userRepository repositories.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}
