package handler

import (
	"add-thanks/internal/constant"
	"add-thanks/internal/gen"
	"add-thanks/internal/usecase/department"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IDepartmentHandler interface {
	PostDepartment(ctx echo.Context) error
}

type DepartmentHandler struct {
	DepartmentUC department.IDepartmentUseCase
}

func NewDepartmentHandler(DepartmentUC department.IDepartmentUseCase) IDepartmentHandler {
	return &DepartmentHandler{DepartmentUC: DepartmentUC}
}

func (dh *DepartmentHandler) PostDepartment(ctx echo.Context) error {
	slog.Info(fmt.Sprintf(constant.Start, constant.PostDepartment))

	var req gen.PostDepartmentJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		res := ErrorResponses(constant.ErrCodeBind, fmt.Sprintf(constant.ErrBind, constant.PostDepartment))
		return ctx.JSON(http.StatusBadRequest, res)
	}
	slog.Info(fmt.Sprintf(constant.End, constant.PostDepartment))
	return ctx.JSON(http.StatusCreated, nil)
}
