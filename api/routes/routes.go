package routes

import (
	"api/config"
	"api/docs"

	"api/app/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(init *config.Initialization) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	docs.SwaggerInfo.BasePath = "/api/"
	router.Static("/public/images/users/", "./public/images/users/")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		auth.POST("/login", init.AuthCtrl.Login)
		auth.GET("/me", middlewares.JWTMiddleware(), init.AuthCtrl.Me)
		auth.GET("/logout", middlewares.JWTMiddleware(), init.AuthCtrl.Logout)
		auth.GET("/refresh", middlewares.JWTMiddleware(), init.AuthCtrl.Refresh)

		user := api.Group("/users")
		user.POST("", middlewares.JWTMiddleware(), init.UserCtrl.AddUser)
		user.GET("", init.UserCtrl.GetUsers)
		user.GET(":id", init.UserCtrl.FindUser)

	}

	router.Run(":8080")
	return router
}
