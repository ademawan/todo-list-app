package entities

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	TaskUid  string `gorm:"index;unique;type:varchar(22)" json:"task_uid"`
	UserUid  string `gorm:"index;type:varchar(22)" json:"-"`
	Title    string `gorm:"type:varchar(30)" json:"title"`
	Priority string `gorm:"type:varchar(20)" json:"priority"`
	Status   string `gorm:"type:varchar(20)" json:"status"`
	Note     string `gorm:"type:varchar(250)" json:"note"`

	CreatedAt      time.Time      `json:"-"`
	UpdatedAt      time.Time      `json:"-"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	Todo_date_time time.Time      `gorm:"index" json:"todo_date_time"`
}
