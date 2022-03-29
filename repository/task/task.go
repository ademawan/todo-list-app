package task

import (
	"errors"
	"todo-list-app/entities"

	"gorm.io/gorm"
)

type TaskRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		database: db,
	}
}

func (tr *TaskRepository) Create(t entities.Task) (entities.Task, error) {
	if err := tr.database.Create(&t).Error; err != nil {
		return t, err
	}

	return t, nil
}

func (tr *TaskRepository) Get(userUid string) ([]entities.Task, error) {
	arrTask := []entities.Task{}
	res := tr.database.Where("user_uid =?", userUid).Find(&arrTask)

	if res.Error != nil {
		return nil, errors.New("failed to get tasks")
	}
	if res.RowsAffected == 0 {
		return arrTask, errors.New("task is empty")
	}

	return arrTask, nil
}

func (tr *TaskRepository) GetByUid(userUid, taskUid string) (entities.Task, error) {
	task := entities.Task{}
	res := tr.database.Where("user_uid =? AND task_uid =?", userUid, taskUid).First(&task, taskUid)

	if res.Error != nil {
		return task, errors.New("failed to get task")
	}
	if res.RowsAffected == 0 {
		return task, errors.New("task not found")
	}

	return task, nil
}

func (tr *TaskRepository) Update(taskUid string, newTask entities.Task) (entities.Task, error) {

	var task entities.Task
	res := tr.database.Where("user_uid =? AND task_uid =?", newTask.UserUid, taskUid).First(&task)
	if res.Error != nil {
		return entities.Task{}, errors.New("failed to update task")
	}
	if res.RowsAffected == 0 {
		return entities.Task{}, errors.New("task not found")
	}

	if err := tr.database.Model(&task).Updates(&newTask).Error; err != nil {
		return entities.Task{}, errors.New("failed to update task")
	}

	return task, nil
}

func (tr *TaskRepository) Delete(userUid, taskUid string) error {
	res := tr.database.Where("user_uid =? AND task_uid =?", userUid, taskUid).Delete(&entities.Task{})
	if res.Error != nil {
		return res.Error
	}
	return nil

}
