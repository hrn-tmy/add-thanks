package user

import "add-thanks/internal/domain/repository"

type IUserUseCase interface{}

type UserUseCase struct {
	userRepo repository.IUserRepository
}

func NewUserUseCase(userRepo repository.IUserRepository) IUserUseCase {
	return &UserUseCase{userRepo: userRepo}
}
