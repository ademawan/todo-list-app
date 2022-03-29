package entities

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Nama     string
	Priority int
	User_ID  int `gorm:"index;column:user_id" json:"user_id"`
	Status   int
	Category string
}
