package auth

import "todo-list-app/entities"

type Auth interface {
	Login(email, password string) (entities.User, error)
}
