package task

import "todo-list-app/entities"

type Task interface {
	Get(userUid string) ([]entities.Task, error)
	GetByUid(userUid, taskId string) (entities.Task, error)
	Create(newTask entities.Task) (entities.Task, error)
	Update(taskUid string, newTask entities.Task) (entities.Task, error)
	Delete(userUid, taskUid string) error
	GetTaskToday(userUid string) ([]entities.Task, error)
}
