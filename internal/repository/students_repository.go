package repository

import (
	"github.com/Jaider-Nieto/Prueba-tecnica-Go/internal/models"
	"gorm.io/gorm"
)

type StudentsRepository struct {
	DB *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentsRepository {
	return &StudentsRepository{
		DB: db,
	}
}

func (r *StudentsRepository) FindAll() ([]models.Students, error) {
	var students []models.Students

	result := r.DB.Find(&students)
	if result.Error != nil {
		return nil, result.Error
	}
	return students, nil
}
func (r *StudentsRepository) FindById(id uint) (*models.Students, error) {
	var student models.Students

	result := r.DB.First(&student, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &student, nil
}
func (r *StudentsRepository) Create(studentBody models.CreateStudent) error {
	student := models.Students{Name: studentBody.Name, Age: studentBody.Age}
	result := r.DB.Create(&student)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *StudentsRepository) Update(studentUpdate models.Students) error {
	result := r.DB.Save(&studentUpdate)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
func (r *StudentsRepository) Delete(id uint) error {
	result := r.DB.Delete(&models.Students{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
