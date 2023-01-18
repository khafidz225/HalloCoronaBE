package models

type User struct {
	ID       int    `json:"id" gorm:"primary_key:auto_increment"`
	Fullname string `json:"fullname" gorm:"type : varchar (255)"`
	Email    string `json:"email" gorm:"type : varchar (255)"`
	Password string `json:"password" gorm:"type : varchar (255)"`
	Gender   string `json:"gender" gorm:"type:varchar(255)"`
	Phone    string `json:"phone" gorm:"type : varchar (255)"`
	Address  string `json:"address" gorm:"type : varchar (255)"`
	Role     string `json:"role" gorm:"type : varchar (255)"`
	Image    string `json:"image" gorm:"type : varchar (255)"`
}

type UserResponse struct {
	ID       int    `json:"id" form:"id"`
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Gender   string `json:"gender" form:"gender"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Role     string `json:"role" form:"role"`
	Image    string `json:"image" form:"image"`
}

func (UserResponse) TableName() string {
	return "users"
}
