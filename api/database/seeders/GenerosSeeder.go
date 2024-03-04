package seeders

import (
	"api/database/migrations"

	"gorm.io/gorm"
)

func GenerosSeeder(DB *gorm.DB) {
	DB.Create(&migrations.Roles{
		Descripcion: "Normal",
	})

	DB.Create(&migrations.Generos{
		Descripcion: "Indefinido",
	})

}
