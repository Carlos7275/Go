package migrations

import (
	"time"
)

type Usuarios struct {
	ID              int    `gorm:"primaryKey"`
	Email           string `gorm:"uniqueIndex"`
	Password        string
	Nombres         string
	ApellidoPaterno string
	ApellidoMaterno string
	Telefono        string
	FechaNacimiento time.Time
	URLImagen       string  `gorm:"default:'/public/images/users/default.png'"`
	IDRol           int     `gorm:"foreignKey:IDRol"`
	Rol             Roles   `gorm:"foreignKey:IDRol"`
	IDGenero        int     `gorm:"foreignKey:IDGenero"`
	Genero          Generos `gorm:"foreignKey:IDGenero"`
	GoogleID        string
	Estatus         int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
