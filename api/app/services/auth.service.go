package services

import (
	"api/app/models/dto"
	"api/app/repositories"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthService interface {
	Login(email string, password string) (string, error)
	Logout(token string)
	Me(token string) (*dto.UsuariosDTO, error)
	RefreshToken(token string) (string, error)
}

type AuthServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewAuthServiceImpl(userRepository repositories.UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepository: userRepository,
	}
}

func (s *AuthServiceImpl) Login(email string, password string) (string, error) {
	user, err := s.userRepository.FindUser(email, password)

	if err != nil {
		return "", err
	}
	userClaims := UserClaims{
		Id: user.ID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	jwt, err := NewAccessToken(userClaims)
	if err != nil {
		log.Fatal("error creating access token")
	}

	return jwt, nil
}

func (s *AuthServiceImpl) Logout(token string) {
	go RevokeToken(token)

}
func (s *AuthServiceImpl) Me(token string) (*dto.UsuariosDTO, error) {
	claims, _ := ExtractClaims(token)

	idfloat := claims["id"].(float64)
	id := int(idfloat)
	user, err := s.userRepository.FindUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthServiceImpl) RefreshToken(token string) (string, error) {
	claims, _ := ExtractClaims(token)
	RevokeToken(token)

	idfloat := claims["id"].(float64)
	id := int(idfloat)
	user, err := s.userRepository.FindUserById(id)
	if err != nil {
		return "", nil
	}

	userClaims := UserClaims{
		Id: user.ID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	jwt, err := NewAccessToken(userClaims)
	if err != nil {
		log.Fatal("error creating access token")
	}

	return jwt, nil

}

func AuthServiceInit(userRepository repositories.UserRepository) *AuthServiceImpl {
	return NewAuthServiceImpl(userRepository)
}
