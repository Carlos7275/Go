package seeders

import (
	"api/database/migrations"

	"gorm.io/gorm"
)

func RolesSeeder(DB *gorm.DB) {

	DB.Create(&migrations.Generos{
		Descripcion: "Masculino",
	})
	DB.Create(&migrations.Generos{
		Descripcion: "Femenino",
	})

	DB.Create(&migrations.Roles{
		Descripcion: "Administrador",
	})

}
