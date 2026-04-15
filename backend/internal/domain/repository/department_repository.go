package repository

import "gorm.io/gorm"

type IDepartmentRepository interface{}

type DepartmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) IDepartmentRepository {
	return &DepartmentRepository{db: db}
}
