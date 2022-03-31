package task

import (
	"errors"
	"time"
	"todo-list-app/entities"

	"github.com/lithammer/shortuuid"

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

func (tr *TaskRepository) Create(task entities.Task) (entities.Task, error) {
	uid := shortuuid.New()
	task.TaskUid = uid
	if err := tr.database.Create(&task).Error; err != nil {
		return task, err
	}

	return task, nil
}

func (tr *TaskRepository) Get(userUid string) ([]entities.Task, error) {
	arrTask := []entities.Task{}
	result := tr.database.Where("user_uid =?", userUid).Find(&arrTask)

	if result.Error != nil {
		return nil, errors.New("failed to get tasks")
	}
	if result.RowsAffected == 0 {
		return arrTask, errors.New("task is empty")
	}

	return arrTask, nil
}
func (tr *TaskRepository) GetTaskToday(userUid string) ([]entities.Task, error) {
	arrTask := []entities.Task{}
	timeNow := time.Now()
	day := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 00, 00, 00, timeNow.Nanosecond(), timeNow.Location())
	endday := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day()+1, 00, 00, 00, timeNow.Nanosecond(), timeNow.Location())
	result := tr.database.Raw("SELECT * FROM tasks WHERE user_uid=? AND todo_date_time BETWEEN ? AND ?", userUid, day, endday).Scan(&arrTask)

	if result.Error != nil {
		return nil, errors.New("failed to get tasks")
	}
	if result.RowsAffected == 0 {
		return arrTask, errors.New("task is empty")
	}

	return arrTask, nil
}

func (tr *TaskRepository) GetByUid(userUid, taskUid string) (entities.Task, error) {
	task := entities.Task{}
	result := tr.database.Where("user_uid =? AND task_uid =?", userUid, taskUid).First(&task)

	if result.Error != nil {
		return task, errors.New("failed to get task")
	}
	if result.RowsAffected == 0 {
		return task, errors.New("task not found")
	}

	return task, nil
}

func (tr *TaskRepository) Update(taskUid string, newTask entities.Task) (entities.Task, error) {

	var task entities.Task
	result := tr.database.Where("user_uid =? AND task_uid =?", newTask.UserUid, taskUid).First(&task)
	if result.Error != nil {
		return entities.Task{}, errors.New("failed to update task")
	}
	if result.RowsAffected == 0 {
		return entities.Task{}, errors.New("task not found")
	}

	if err := tr.database.Model(&task).Where("user_uid=? AND task_uid =?", newTask.UserUid, taskUid).Updates(&newTask).Error; err != nil {
		return entities.Task{}, errors.New("failed to update task")
	}

	return task, nil
}

func (tr *TaskRepository) Delete(userUid, taskUid string) error {
	result := tr.database.Where("user_uid =? AND task_uid =?", userUid, taskUid).Delete(&entities.Task{})
	if result.Error != nil {
		return result.Error
	}
	return nil

}
