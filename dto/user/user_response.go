package usersdto

type UserResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
}

type UpdateUserRequest struct {
	// FullName string `json:"fullname" form:"fullname"`
	// Email    string `json:"email" form:"email"`
	// Password string `json:"password" form:"password"`
	// Phone    string `json:"phone" form:"phone" `
	// Address  string `json:"address" form:"address" `
	// Role     string `json:"role" form:"role"`
	Image string `json:"image" form:"image"`
}

type UserResponseDel struct {
	ID int `json:"id"`
}
