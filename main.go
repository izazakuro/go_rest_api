package main

import (
	"go_rest_api/controller"
	"go_rest_api/db"
	"go_rest_api/repository"
	"go_rest_api/router"
	"go_rest_api/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
