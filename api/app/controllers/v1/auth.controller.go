package v1

import (
	"api/app/models/requests"
	"api/app/services"
	"api/pkg"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
	Me(c *gin.Context)
	Logout(c *gin.Context)
	Refresh(c *gin.Context)
}

type AuthControllerImpl struct {
	svc services.AuthService
}

//	@BasePath	/api/v1

// Iniciar Sesión
//	@Summary	Login
//	@Tags		Auth
//	@Accept		json
//	@Param		Login	body	requests.LoginRequest	true	"login  request"
//	@Produce	json
//	@Success	200	{string}	token
//	@Router		/v1/auth/login/ [Post]
func (ac *AuthControllerImpl) Login(c *gin.Context) {
	var loginRequest requests.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse__(err.Error()))
		return
	}

	response, err := ac.svc.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse__(err.Error()))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse_("Operacion Exitosa", response))
}

//	@Summary	Me
//	@Tags		Auth
//	@Produce	json
//	@Router		/v1/auth/me/ [Get]
// @Security Bearer
func (ac *AuthControllerImpl) Me(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	user, err := ac.svc.Me(tokenString)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, pkg.BuildResponse_("Operacion Exitosa", user))
}

func (ac *AuthControllerImpl) Logout(c *gin.Context) {
	ac.Logout(c)
	c.JSON(http.StatusOK, pkg.BuildResponse__("Se cerro sesion con exito"))

}

func (ac *AuthControllerImpl) Refresh(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	token, err := ac.svc.RefreshToken(tokenString)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, pkg.BuildResponse_("Operacion Exitosa", token))

}

func AuthControllerInit(authService services.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		svc: authService,
	}
}
