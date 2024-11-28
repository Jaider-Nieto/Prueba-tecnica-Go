package services

import (
	"errors"

	"github.com/Jaider-Nieto/Prueba-tecnica-Go/internal/models"
	"github.com/Jaider-Nieto/Prueba-tecnica-Go/internal/repository"
)

type StudentService struct {
	reposository *repository.StudentsRepository
}

func NewStudentService(repository *repository.StudentsRepository) *StudentService {
	return &StudentService{
		reposository: repository,
	}
}

func (s *StudentService) GetAll() ([]models.Students, error) {
	return s.reposository.FindAll()
}
func (s *StudentService) GetOne(id uint) (*models.Students, error) {
	return s.reposository.FindById(id)
}
func (s *StudentService) Create(student models.CreateStudent) error {
	return s.reposository.Create(student)
}
func (s *StudentService) Update(studentUpdate models.UpdateStudent) error {
	student, err := s.reposository.FindById(studentUpdate.ID)
	if err != nil {
		return err
	}
	if student.ID == 0 {
		return errors.New("student not found")
	}

	if studentUpdate.Age != nil {
		student.Age = *studentUpdate.Age
	}
	if studentUpdate.Name != nil {
		student.Name = *studentUpdate.Name
	}

	return s.reposository.Update(*student)
}
func (s *StudentService) Delete(id uint) error {
	return s.reposository.Delete(id)
}
