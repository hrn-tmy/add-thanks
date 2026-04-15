package handler

import "add-thanks/internal/usecase/user"

type IUserHandler interface{}

type UserHandler struct {
	userUC user.IUserUseCase
}

func NewUserHandler(userUC user.IUserUseCase) IUserHandler {
	return &UserHandler{userUC: userUC}
}
