package main

import (
	"add-thanks/internal/domain/repository"
	"add-thanks/internal/handler"
	"add-thanks/internal/infra/database"
	"add-thanks/internal/logger"
	"add-thanks/internal/router"
	"add-thanks/internal/usecase/department"
	"add-thanks/internal/usecase/user"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	logger.NewLogger()
	db, err := database.NewDB()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	userRepo := repository.NewUserRepository(db)
	departmentRepo := repository.NewDepartmentRepository(db)

	userUC := user.NewUserUseCase(userRepo)
	departmentUC := department.NewDepartmentUseCase(departmentRepo)

	userHandler := handler.NewUserHandler(userUC)
	departmentHandler := handler.NewDepartmentHandler(departmentUC)

	e := echo.New()

	router.NewRouter(e, userHandler, departmentHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
