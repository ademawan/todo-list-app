package main

import (
	"fmt"
	"todo-list-app/configs"
	ac "todo-list-app/delivery/controllers/auth"
	tc "todo-list-app/delivery/controllers/task"
	uc "todo-list-app/delivery/controllers/user"
	"todo-list-app/delivery/routes"
	authRepo "todo-list-app/repository/auth"
	taskRepo "todo-list-app/repository/task"
	userRepo "todo-list-app/repository/user"
	"todo-list-app/utils"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	"github.com/labstack/gommon/log"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

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
	e.Validator = &CustomValidator{validator: validator.New()}

	routes.RegisterPath(e, authController, userController, taskController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
