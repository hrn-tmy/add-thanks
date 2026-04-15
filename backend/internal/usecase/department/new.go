package department

import "add-thanks/internal/domain/repository"

type IDepartmentUseCase interface{}

type DepartmentUseCase struct {
	DepartmentRepo repository.IDepartmentRepository
}

func NewDepartmentUseCase(DepartmentRepo repository.IDepartmentRepository) IDepartmentUseCase {
	return &DepartmentUseCase{DepartmentRepo: DepartmentRepo}
}
