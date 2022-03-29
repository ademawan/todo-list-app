package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserUid  string `gorm:"index;unique;type:varchar(22)" json:"user_uid"`
	Name     string `gorm:"type:varchar(30)" json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `gorm:"-" json:"-"`
	Address  string `gorm:"type:varchar(200)" json:"address"`
	Gender   string `gorm:"type:enum('male','female')" json:"gender"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `gorm:"-" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Task []Task `gorm:"ForeignKey:UserUid"`
}
