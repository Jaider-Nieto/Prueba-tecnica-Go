package controllers

import (
	"net/http"
	"strconv"

	"github.com/Jaider-Nieto/Prueba-tecnica-Go/internal/models"
	"github.com/Jaider-Nieto/Prueba-tecnica-Go/internal/services"
	"github.com/gin-gonic/gin"
)

type StudentsController struct {
	service *services.StudentService
}

func NewStudentsController(service *services.StudentService) *StudentsController {
	return &StudentsController{
		service: service,
	}
}

func (ctr *StudentsController) GetStudents(c *gin.Context) {
	students, err := ctr.service.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, students)
}
func (ctr *StudentsController) GetStudent(c *gin.Context) {
	id, err := (strconv.ParseUint(c.Param("id"), 10, 0))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Id invalid")
		return
	}
	students, err := ctr.service.GetOne(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, students)
}
func (ctr *StudentsController) PostStudent(c *gin.Context) {
	var student models.CreateStudent

	if err := c.BindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if err := ctr.service.Create(student); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, "created")
}
func (ctr *StudentsController) PatchStudent(c *gin.Context) {
	var student models.UpdateStudent

	if err := c.BindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if err := ctr.service.Update(student); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, "created")
}
func (ctr *StudentsController) DeleteStudent(c *gin.Context) {
	id, errId := (strconv.ParseUint(c.Param("id"), 10, 0))
	if errId != nil {
		c.JSON(http.StatusBadRequest, "Id invalid")
		return
	}
	_, err := ctr.service.GetOne(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := ctr.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, "deletd")
}
