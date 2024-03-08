package dto

import (
	"api/database/migrations"
	"time"
)

type UsuariosDTO struct {
	ID              int
	Email           string
	Nombres         string
	ApellidoPaterno string
	ApellidoMaterno string
	Telefono        string
	FechaNacimiento time.Time
	URLImagen       string
	Rol             migrations.Roles
	Genero          migrations.Generos
	Estatus         int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
