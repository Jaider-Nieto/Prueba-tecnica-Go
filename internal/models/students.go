package models

type Students struct {
	ID             uint             `gorm:"primaryKey"`
	Name           string           `gorm:"not null"`
	Age            uint             `gorm:"not null"`
}

type CreateStudent struct {
	Name string
	Age  uint
}

type UpdateStudent struct {
	ID uint
	Name *string
	Age  *uint
}
