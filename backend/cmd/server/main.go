package main

import (
	"add-thanks/internal/domain/repository"
	"add-thanks/internal/handler"
	"add-thanks/internal/infra/database"
	"add-thanks/internal/router"
	"add-thanks/internal/usecase/user"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := database.NewDB()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	userRepo := repository.NewUserRepository(db)
	userUC := user.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUC)

	e := echo.New()

	router.NewRouter(e, userHandler)

	e.Logger.Fatal(e.Start(":8080"))
}