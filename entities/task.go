package entities

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	TaskUid  string `gorm:"index;unique;type:varchar(22)" json:"task_uid"`
	UserUid  string `gorm:"index;type:varchar(22)" json:"user_uid"`
	Title    string `gorm:"type:varchar(30)" json:"title"`
	Priority string `gorm:"type:enum('hight','medium','low')" json:"gender"`
	User_ID  int    `gorm:"index;column:user_id" json:"user_id"`
	Status   int    `gorm:"type:enum('done','undone','ignore')" json:"status"`
	Note     string `gorm:"type:varchar(250)" json:"note"`

	CreatedAt      time.Time      `json:"-"`
	UpdatedAt      time.Time      `json:"-"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	Todo_date_time time.Time      `gorm:"index" json:"todo_date_time"`
}
