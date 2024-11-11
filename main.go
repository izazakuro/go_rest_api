package main

import (
	"go_rest_api/controller"
	"go_rest_api/db"
	"go_rest_api/repository"
	"go_rest_api/router"
	"go_rest_api/usecase"
	"go_rest_api/validator"
)

func main() {
	db := db.NewDB()

	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()

	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)

	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)

	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
