package main

import (
	"api/app/models/dto"
	"api/config"
	"api/database/migrations"
	"api/routes"

	"github.com/devfeel/mapper"
)

//	@title			Swagger Golang Api
//	@description	Ejemplo de API en Go Hecha por Carlos Sandoval.
//	@BasePath		/api
//	@versionList
//
//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
func main() {
	mapper.Register(&migrations.Usuarios{})
	mapper.Register(&dto.UsuariosDTO{})
	init := config.Init()
	config.InitLog()
	routes.Init(init)
}
