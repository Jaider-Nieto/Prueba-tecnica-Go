package interfaces

type StudentsControllerInterfaces interface {
	GetAllStudents()
	GetOneStudentById()
	PostStudent()
	PatchStudent()
	DeleteStudent()
}