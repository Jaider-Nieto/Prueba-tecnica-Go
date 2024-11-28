package routes

import (
	"github.com/Jaider-Nieto/Prueba-tecnica-Go/internal/config"
	"github.com/Jaider-Nieto/Prueba-tecnica-Go/internal/controllers"
	"github.com/Jaider-Nieto/Prueba-tecnica-Go/internal/repository"
	"github.com/Jaider-Nieto/Prueba-tecnica-Go/internal/services"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {

	db := config.DB
	r := repository.NewStudentRepository(db)
	s := services.NewStudentService(r)
	c := controllers.NewStudentsController(s)

	students := router.Group("students")
	{
		students.GET("/all", c.GetStudents)
		students.GET("/id/:id", c.GetStudent)
		students.POST("/create", c.PostStudent)
		students.PATCH("/update", c.PatchStudent)
		students.DELETE("/delete/:id", c.DeleteStudent)
	}

}
