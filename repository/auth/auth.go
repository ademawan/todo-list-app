package auth

import (
	"todo-list-app/entities"
	"todo-list-app/middlewares"

	"errors"

	"gorm.io/gorm"
)

type AuthDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AuthDb {
	return &AuthDb{
		db: db,
	}
}

func (ad *AuthDb) Login(email, password string) (entities.User, error) {
	user := entities.User{}

	if err := ad.db.Model(&user).Where("email = ?", email).First(&user).Error; err != nil {
		return user, errors.New("email not found")
	}

	match := middlewares.CheckPasswordHash(password, user.Password)

	if !match {
		return user, errors.New("incorrect password")
	}

	return user, nil
}
