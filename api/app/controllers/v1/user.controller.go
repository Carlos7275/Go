package v1

import (
	"api/app/services"
	"api/pkg"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController interface {
	AddUser(c *gin.Context)
	FindUser(c *gin.Context)
	GetUsers(c *gin.Context)
	ChangeUserStatus(c *gin.Context)
}

type UserControllerImpl struct {
	svc services.UserService
}

// Add User
//
//	@Summary	Add User
//	@Tags		Users
//	@Param		User	body	requests.UserRequest	true	"User Request"
//	@Accept		json
//	@Produce	json
//	@Success	200	{string}	string
//	@Failure	400	{string}	string
//	@Failure	401	{string}	string
//	@Failure	500	{string}	string
//
//	@Router		/v1/users/ [POST]
//
//	@Security	Bearer
func (uc *UserControllerImpl) AddUser(c *gin.Context) {

}

// Get Users
//
//	@Summary	Get User
//	@Tags		Users
//	@Accept		json
//
// @Param        id   path      int  true  "User Id"
//
//	@Produce	json
//	@Success	200	{object}	dto.UsuariosDTO
//	@Failure	400	{string}	string
//	@Failure	401	{string}	string
//	@Failure	500	{string}	string
//
//	@Router		/v1/users/{id} [GET]
//
//	@Security	Bearer
func (uc *UserControllerImpl) FindUser(c *gin.Context) {
	strid := c.Param("id")
	id, _ := strconv.Atoi(strid)

	user, err := uc.svc.FindUser(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, pkg.BuildResponse__("No se encontro el usuario"))
		} else {
			c.JSON(http.StatusInternalServerError, pkg.BuildResponse__(err.Error()))
		}
		return
	}
	c.JSON(http.StatusOK, pkg.BuildResponse_("Operacion Exitosa", user))
}

// Get Users
//
//	@Summary	Get all  Users Registered
//	@Tags		Users
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	[]dto.UsuariosDTO
//	@Failure	400	{string}	string
//	@Failure	401	{string}	string
//	@Failure	500	{string}	string
//
//	@Router		/v1/users/ [GET]
//
//	@Security	Bearer
func (uc *UserControllerImpl) GetUsers(c *gin.Context) {
	users, err := uc.svc.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse__(err.Error()))
	}
	c.JSON(http.StatusOK, pkg.BuildResponse_("Operacion exitosa", users))
}

func (uc *UserControllerImpl) ChangeUserStatus(c *gin.Context) {

}

func UserControllerInit(userService services.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		svc: userService,
	}
}
