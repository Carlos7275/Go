package seeders

import (
	"api/database/migrations"
	"time"

	"github.com/go-crypt/crypt/algorithm"
	"github.com/go-crypt/crypt/algorithm/argon2"
	"gorm.io/gorm"
)

func UserSeeder(DB *gorm.DB) {

	var (
		hasher *argon2.Hasher
		digest algorithm.Digest
		err    error
	)

	if hasher, err = argon2.New(
		argon2.WithProfileRFC9106LowMemory(),
	); err != nil {
		panic(err)
	}

	if digest, err = hasher.Hash("admin"); err != nil {
		panic(err)
	}

	DB.Create(&migrations.Usuarios{
		Email:           "admin@admin.com",
		Password:        digest.Encode(),
		Nombres:         "Carlos Fernando",
		ApellidoMaterno: "Sandoval",
		ApellidoPaterno: "Lizárraga",
		Telefono:        "6682566496",
		FechaNacimiento: time.Date(2001, 2, 14, 0, 0, 0, 0, time.Local),
		IDRol:           1,
		IDGenero:        1,
		Estatus:         1,
	})

}
