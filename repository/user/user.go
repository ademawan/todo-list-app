package user

import (
	"errors"
	"todo-list-app/entities"
	"todo-list-app/middlewares"

	"github.com/lithammer/shortuuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (ur *UserRepository) Register(u entities.User) (entities.User, error) {

	u.Password, _ = middlewares.HashPassword(u.Password)
	uid := shortuuid.New()
	u.UserUid = uid

	if err := ur.database.Create(&u).Error; err != nil {
		return u, errors.New("invalid input or this email was created (duplicated entry)")
	}

	return u, nil
}

func (ur *UserRepository) GetByUid(user_uid string) (entities.User, error) {
	arrUser := entities.User{}

	result := ur.database.Preload("Task").Where("user_uid =?", user_uid).First(&arrUser)
	if result.RowsAffected == 0 {
		return arrUser, errors.New("record not found")
	}
	if err := result.Error; err != nil {
		return arrUser, err
	}

	return arrUser, nil
}

func (ur *UserRepository) Update(user_uid string, newUser entities.User) (entities.User, error) {

	var user entities.User
	ur.database.Where("user_uid =?", user_uid).First(&user)

	if err := ur.database.Model(&user).Where("user_uid =?", user_uid).Updates(&newUser).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) Delete(user_uid string) error {

	if err := ur.database.Where("user_uid = ?", user_uid).Delete(&entities.User{}).Error; err != nil {
		return err
	}
	return nil

}
