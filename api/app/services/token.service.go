package services

import (
	"errors"
	"os"
	"sync"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

var (
	blacklistTokens = make(map[string]bool)
	mutex           = sync.RWMutex{}
)

func RevokeToken(tokenString string) {
	mutex.Lock()
	defer mutex.Unlock()
	blacklistTokens[tokenString] = true
}

func IsTokenRevoked(tokenString string) bool {
	mutex.RLock()
	defer mutex.RUnlock()
	return blacklistTokens[tokenString]
}

func NewAccessToken(claims UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

func ExtractClaims(tokenString string) (jwt.MapClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("token claims are not of type jwt.MapClaims")
	}

	return claims, nil
}

func NewRefreshToken(claims jwt.StandardClaims) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return refreshToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

func ParseAccessToken(accessToken string) *UserClaims {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	return parsedAccessToken.Claims.(*UserClaims)
}

func ParseRefreshToken(refreshToken string) *jwt.StandardClaims {
	parsedRefreshToken, _ := jwt.ParseWithClaims(refreshToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	return parsedRefreshToken.Claims.(*jwt.StandardClaims)
}
