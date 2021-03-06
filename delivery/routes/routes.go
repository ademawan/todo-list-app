package routes

import (
	"todo-list-app/delivery/controllers/auth"
	"todo-list-app/delivery/controllers/task"

	"todo-list-app/delivery/controllers/user"
	"todo-list-app/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo,
	aa *auth.AuthController,
	uc *user.UserController,
	tc *task.TaskController,

) {

	//CORS
	e.Use(middleware.CORS())

	//LOGGER
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	//ROUTE REGISTER - LOGIN USERS
	e.POST("users/register", uc.Register())
	e.POST("users/login", aa.Login())
	e.POST("users/logout", aa.Logout(), middlewares.JwtMiddleware())

	//ROUTE USERS
	e.GET("/users/me", uc.GetByUid(), middlewares.JwtMiddleware())
	e.PUT("/users/me", uc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/users/me", uc.Delete(), middlewares.JwtMiddleware())
	//ROUTE Task
	e.POST("/users/me/tasks", tc.Create(), middlewares.JwtMiddleware())
	e.GET("/users/me/tasks", tc.Get(), middlewares.JwtMiddleware())
	e.GET("/users/me/tasks/today", tc.GetTaskToday(), middlewares.JwtMiddleware())
	e.GET("/users/me/tasks/:task_uid", tc.GetByUid(), middlewares.JwtMiddleware())
	e.PUT("/users/me/tasks/:task_uid", tc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/users/me/tasks/:task_uid", tc.Delete(), middlewares.JwtMiddleware())
}
