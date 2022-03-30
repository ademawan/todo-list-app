package task

import (
	"net/http"
	"todo-list-app/delivery/controllers/common"
	"todo-list-app/entities"
	"todo-list-app/middlewares"
	"todo-list-app/repository/task"

	"github.com/labstack/echo/v4"
)

type TaskController struct {
	repo task.Task
}

func New(repository task.Task) *TaskController {
	return &TaskController{
		repo: repository,
	}
}
func (tc *TaskController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		task := TaskRequestFormat{}
		userUid := middlewares.ExtractTokenUserUid(c)
		c.Bind(&task)

		err := c.Validate(&task)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ResponseUser(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := tc.repo.Create(entities.Task{
			UserUid:        userUid,
			Title:          task.Title,
			Priority:       task.Priority,
			Note:           task.Note,
			Todo_date_time: task.Todo_date_time,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusCreated, common.ResponseUser(http.StatusCreated, "Success create task", res))
	}
}

func (tc *TaskController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		userUid := middlewares.ExtractTokenUserUid(c)

		res, err := tc.repo.Get(userUid)

		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "task is empty" {
				statusCode = http.StatusOK
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(statusCode, errorMessage, nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success get all task", res))
	}
}

func (tc *TaskController) GetByUid() echo.HandlerFunc {
	return func(c echo.Context) error {
		taskUid := c.Param("task_uid")
		userUid := middlewares.ExtractTokenUserUid(c)

		res, err := tc.repo.GetByUid(userUid, taskUid)

		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "task not found" {
				statusCode = http.StatusNotFound
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(http.StatusNotFound, errorMessage, nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success get task by uid", res))
	}
}

func (tc *TaskController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newTask = UpdateTaskRequestFormat{}
		taskUid := c.Param("task_uid")
		userUid := middlewares.ExtractTokenUserUid(c)
		c.Bind(&newTask)

		err := c.Validate(&newTask)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ResponseUser(http.StatusBadRequest, "There is some problem from input", nil))
		}
		res, err := tc.repo.Update(taskUid, entities.Task{
			UserUid:        userUid,
			Title:          newTask.Title,
			Priority:       newTask.Priority,
			Status:         newTask.Status,
			Note:           newTask.Note,
			Todo_date_time: newTask.Todo_date_time,
		})

		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "task not found" {
				statusCode = http.StatusNotFound
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(http.StatusNotFound, errorMessage, nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success update task", res))
	}
}

func (tc *TaskController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		taskUid := c.Param("task_uid")
		userUid := middlewares.ExtractTokenUserUid(c)

		err := tc.repo.Delete(userUid, taskUid)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.ResponseUser(http.StatusNotFound, "not found", nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success delete task", nil))
	}
}
