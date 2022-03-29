package user

import "todo-list-app/entities"

type User interface {
	Register(user entities.User) (entities.User, error)
	GetByUid(userUid string) (entities.User, error)
	Update(userUid string, newUser entities.User) (entities.User, error)
	Delete(userUid string) error
	//jika diperlukan
	// GetAll() ([]entities.User, error)
}
