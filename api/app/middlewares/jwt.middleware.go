package middlewares

import (
	"api/app/services"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is missing"})
			c.Abort()
			return
		}
		tokenSecret := []byte(os.Getenv("TOKEN_SECRET"))

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return tokenSecret, nil
		})

		if err != nil || !token.Valid || services.IsTokenRevoked(tokenString) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
