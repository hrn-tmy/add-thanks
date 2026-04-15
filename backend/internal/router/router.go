package router

import (
	"add-thanks/internal/gen"
	"add-thanks/internal/handler"

	"github.com/labstack/echo/v4"
)

type ServerHandler struct {
	userHandler       handler.IUserHandler
	departmentHandler handler.IDepartmentHandler
}

func (h *ServerHandler) PostDepartment(ctx echo.Context) error {
	return h.departmentHandler.PostDepartment(ctx)
}

func NewRouter(e *echo.Echo, userHandler handler.IUserHandler, departmentHandler handler.IDepartmentHandler) {
	handler := &ServerHandler{
		userHandler:       userHandler,
		departmentHandler: departmentHandler,
	}
	wrapper := gen.ServerInterfaceWrapper{
		Handler: handler,
	}

	dep := e.Group("/departments")
	dep.POST("", wrapper.PostDepartment)
}
