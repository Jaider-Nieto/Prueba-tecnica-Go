package main

import (
	"github.com/Jaider-Nieto/Prueba-tecnica-Go/internal/config"
	"github.com/Jaider-Nieto/Prueba-tecnica-Go/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.DbConnection()
	config.AutoMigrate()

	routes.InitRoutes(r)
	
	r.Run()
}
