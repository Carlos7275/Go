package config

import (
	"api/database/migrations"
	"api/database/seeders"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection() *gorm.DB {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	// Obtener valores de las variables de entorno
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PWD")
	HOST := os.Getenv("HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_NAME")
	MIGRATE := os.Getenv("MIGRATE")
	SEEDING := os.Getenv("SEEDING")
	// Crear cadena de conexión para la base de datos principal
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", HOST, USER, PASS, DBNAME, PORT)

	// Establecer la conexión a la base de datos principal
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	migrate, err := strconv.ParseBool(MIGRATE)
	if err != nil {
		panic(err.Error())
	}
	if migrate {
		db.AutoMigrate(&migrations.Roles{}, &migrations.Generos{}, &migrations.Usuarios{})
	}

	seeding, err := strconv.ParseBool(SEEDING)

	if err != nil {
		panic(err.Error())
	}

	if seeding {
		seeders.RolesSeeder(db)
		seeders.GenerosSeeder(db)
		seeders.UserSeeder(db)
	}
	return db
}
