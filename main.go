package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
)


func main() {
	db := db.NewDB()
	
	userRepository := repository.NewUserRepository(db)
	taskRepogitory := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUsercase(userRepository)
	taskUsecase := usecase.NewTaskUsecase(taskRepogitory)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
