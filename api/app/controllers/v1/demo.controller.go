package v1

import (
	"api/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@BasePath	/api/v1

// PingExample godoc
//	@Summary	ping example
//	@Schemes
//	@Description	do ping
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	Helloworld
//	@Router			/v1/hw/ [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, pkg.BuildResponse__("hello World"))
}
