package repositories

import (
	"api/app/models/dto"
	"api/database/migrations"
	"errors"

	"github.com/devfeel/mapper"
	"github.com/go-crypt/crypt"
	log "github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUser(email string, password string) (dto.UsuariosDTO, error)
	FindAllUser() ([]dto.UsuariosDTO, error)
	FindUserById(id int) (dto.UsuariosDTO, error)
	Save(user *dto.UsuariosDTO) (dto.UsuariosDTO, error)
	DeleteUserById(id int) error
}
type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u UserRepositoryImpl) FindUser(email string, password string) (dto.UsuariosDTO, error) {
	var user = migrations.Usuarios{}

	var err = u.db.Preload("Genero").Preload("Rol").Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// No se encontró ningún registro con el correo electrónico dado
			return dto.UsuariosDTO{}, errors.New("¡Verifique sus Datos!")
		}
		return dto.UsuariosDTO{}, errors.New(err.Error())
	}

	valid, err := crypt.CheckPassword(password, user.Password)
	if err != nil {
		return dto.UsuariosDTO{}, errors.New("Error al verificar la contraseña")
	}

	if !valid {
		return dto.UsuariosDTO{}, errors.New("¡Verifique su Contraseña!")
	}
	var userMap dto.UsuariosDTO

	mapper.Mapper(&user, &userMap)
	return userMap, nil
}

func (u UserRepositoryImpl) FindAllUser() ([]dto.UsuariosDTO, error) {
	var users []migrations.Usuarios

	var err = u.db.Preload("Genero").Preload("Rol").Find(&users).Error
	if err != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	var usersMap []dto.UsuariosDTO
	for _, user := range users {
		var userDTO dto.UsuariosDTO
		err = mapper.Mapper(&user, &userDTO)
		if err != nil {
			log.Error("Error mapping user to DTO: ", err)
			return nil, err
		}
		usersMap = append(usersMap, userDTO)
	}

	return usersMap, nil
}

func (u UserRepositoryImpl) FindUserById(id int) (dto.UsuariosDTO, error) {
	user := migrations.Usuarios{
		ID: id,
	}
	err := u.db.Preload("Genero").Preload("Rol").First(&user).Error
	if err != nil {
		log.Error("Got and error when find user by id. Error: ", err)
		return dto.UsuariosDTO{}, err
	}
	var userMap dto.UsuariosDTO

	mapper.Mapper(&user, &userMap)
	return userMap, nil
}

func (u UserRepositoryImpl) Save(user *dto.UsuariosDTO) (dto.UsuariosDTO, error) {
	var err = u.db.Save(user).Error
	if err != nil {
		log.Error("Got an error when save user. Error: ", err)
		return dto.UsuariosDTO{}, err
	}
	var userMap dto.UsuariosDTO

	mapper.Mapper(&user, &userMap)
	return *&userMap, nil
}

func (u UserRepositoryImpl) DeleteUserById(id int) error {
	err := u.db.Delete(&dto.UsuariosDTO{}, id).Error
	if err != nil {
		log.Error("Got an error when delete user. Error: ", err)
		return err
	}
	return nil
}

func UserRepositoryInit(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}
