package routes

import (
	"api/app/services"
	"api/config"
	"api/docs"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func jwtMiddleware() gin.HandlerFunc {
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

func Init(init *config.Initialization) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	docs.SwaggerInfo.BasePath = "/api/"
	router.Static("/public/images/users/", "./public/images/users/")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	{
		user := api.Group("/auth")
		user.POST("/login", init.AuthCtrl.Login)
		user.GET("/me", jwtMiddleware(), init.AuthCtrl.Me)
		user.GET("/logout", jwtMiddleware(), init.AuthCtrl.Logout)
		user.GET("/refresh", jwtMiddleware(), init.AuthCtrl.Refresh)

	}
	router.Run(":8080")
	return router
}
