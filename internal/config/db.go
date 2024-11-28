package config

import (
	"log"

	"github.com/Jaider-Nieto/Prueba-tecnica-Go/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConnection() {
	var err error

	DB, err = gorm.Open(postgres.Open("host=localhost user=root password=12345 dbname=db port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func AutoMigrate(){
	err := DB.AutoMigrate(
		&models.Students{},
	)
	if err != nil{
		log.Fatal(err)
	}
}
