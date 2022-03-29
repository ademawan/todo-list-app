package task

import (
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

func (tr *TaskRepository) Get() ([]entities.Task, error) {
	arrTask := []entities.Task{}

	if err := tr.database.Find(&arrTask).Error; err != nil {
		return nil, err
	}

	return arrTask, nil
}

func (tr *TaskRepository) GetByUid(taskUid string) (entities.Task, error) {
	arrTask := entities.Task{}

	if err := tr.database.First(&arrTask, taskUid).Error; err != nil {
		return arrTask, err
	}

	return arrTask, nil
}

func (tr *TaskRepository) Update(taskUid string, newTask entities.Task) (entities.Task, error) {

	var task entities.Task
	tr.database.First(&task, taskUid)

	if err := tr.database.Model(&task).Updates(&newTask).Error; err != nil {
		return task, err
	}

	return task, nil
}

func (tr *TaskRepository) Delete(taskUid string) error {

	var task entities.Task

	if err := tr.database.First(&task, taskUid).Error; err != nil {
		return err
	}
	tr.database.Delete(&task, taskUid)
	return nil

}
