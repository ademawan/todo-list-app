package task

import (
	"time"
	"todo-list-app/entities"
)

//----------------------------------------------------
//REQUEST FORMAT
//----------------------------------------------------
type TaskRequestFormat struct {
	UserUid        string
	Title          string    `json:"title" form:"title" validate:"required,min=3,max=20,excludesall=!@#?^#*()_+-=0123456789%&"`
	Priority       string    `json:"priority" form:"priority" validate:"required,min=3,max=20,excludesall=!@#?^#*()_+-=0123456789%&"`
	Note           string    `json:"note" form:"note"`
	Todo_date_time time.Time `json:"todo_date_time" form:"todo_date_time"`
}

type UpdateTaskRequestFormat struct {
	UserUid        string
	Title          string    `json:"title" form:"title"  validate:"omitempty,min=3,max=50,excludesall=!@#?^#*()_+-=0123456789%&"`
	Priority       string    `json:"priority" form:"priority"  validate:"omitempty,min=3,max=6,excludesall=!@#?^#*()_+-=0123456789%&"`
	Status         string    `json:"status" form:"status"  validate:"omitempty,min=4,max=8,excludesall=!@#?^#*()_+-=0123456789%&"`
	Note           string    `json:"note" form:"note"`
	Todo_date_time time.Time `json:"todo_date_time" form:"todo_date_time"`
}

//-----------------------------------------------------
//RESPONSE FORMAT
//-----------------------------------------------------

type GetTasksResponseFormat struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    []entities.Task `json:"data"`
}

type GetTaskResponseFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.Task `json:"data"`
}

type UpdateResponseFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.Task `json:"data"`
}

type DeleteResponseFormat struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ReopenTaskResponFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.Task `json:"data"`
}

type CompleteTaskResponFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.Task `json:"data"`
}
