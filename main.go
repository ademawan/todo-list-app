package main

import (
	"fmt"
	"todo-list-app/configs"
	ac "todo-list-app/delivery/controllers/auth"
	"todo-list-app/delivery/controllers/routes"
	tc "todo-list-app/delivery/controllers/task"
	uc "todo-list-app/delivery/controllers/user"
	authRepo "todo-list-app/repository/auth"
	taskRepo "todo-list-app/repository/task"
	userRepo "todo-list-app/repository/user"
	"todo-list-app/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	authRepo := authRepo.New(db)
	userRepo := userRepo.New(db)
	taskRepo := taskRepo.New(db)

	authController := ac.New(authRepo)
	userController := uc.New(userRepo)
	taskController := tc.New(taskRepo)

	e := echo.New()

	routes.RegisterPath(e, authController, userController, taskController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
