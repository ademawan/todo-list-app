package user

// "gorm.io/gorm"

type UserCreateResponse struct {
	User_uid string `json:"user_uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Gender   string `json:"gender" form:"gender"`
	// Roles    bool   `json:"roles" form:"roles"`
	// Image    string `json:"image" form:"image"`
}
type UserUpdateResponse struct {
	User_uid string `json:"user_uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Gender   string `json:"gender" form:"gender"`
	// Roles    bool   `json:"roles" form:"roles"`
	// Image    string `json:"image" form:"image"`
}
type UserGetByIdResponse struct {
	User_uid string `json:"user_uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Gender   string `json:"gender" form:"gender"`
	// Roles    bool   `json:"roles" form:"roles"`
	// Image    string `json:"image" form:"image"`
}

//=========================================================

// =================== Create User Request =======================
type CreateUserRequestFormat struct {
	Name     string `json:"name" form:"name" validate:"required,min=3,max=20,excludesall=!@#?^#*()_+-=0123456789%&"`
	User_uid string
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=3,max=15"`
	Address  string `json:"address" form:"address" validate:"required"`
	Gender   string `json:"gender" form:"gender" validate:"required"`
}

// =================== Update User Request =======================
type UpdateUserRequestFormat struct {
	Name     string `json:"name" form:"name" validate:"omitempty,min=3,max=20,excludesall=!@#?^#*()_+-=0123456789%&"`
	Email    string `json:"email" form:"email" validate:"omitempty,email"`
	Password string `json:"password" form:"password" validate:"omitempty,required,min=3,max=15"`
	Address  string `json:"address" form:"address" validate:"omitempty"`
	Gender   string `json:"gender" form:"gender"  validate:"omitempty"`
}
