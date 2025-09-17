package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"go-rest-api/validator"
)


func main() {
	db := db.NewDB()

	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	
	userRepository := repository.NewUserRepository(db)
	taskRepogitory := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUsercase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepogitory, taskValidator)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
