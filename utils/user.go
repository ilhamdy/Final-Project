package utils

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type RegisterUser struct {
	Age      uint   `json:"age" example:"8"`
	Username string `json:"username" example:"johndoe"`
	Password string `json:"password" example:"secret"`
	Email    string `json:"email" example:"johndoe@example.com"`
}

type RegisteredUser struct {
	Age      uint   `json:"age" example:"8"`
	Email    string `json:"email" example:"johndoe@example.com"`
	ID       string `json:"id" example:"the user id generated here"`
	Username string `json:"username" example:"johndoe"`
}

type ResponseDataRegisteredUser struct {
	Status string         `json:"status" example:"success"`
	Data   RegisteredUser `json:"data"`
}

type LoginUser struct {
	Email    string `json:"email" example:"johndoe@example.com"`
	Password string `json:"password" example:"secret"`
}

type LoggedinUser struct {
	Token string `json:"token" example:"the token generated here"`
}

type ResponseDataLoggedinUser struct {
	Status string       `json:"status" example:"success"`
	Data   LoggedinUser `json:"data"`
}

type ResponseMessageUser struct {
	Status string `json:"status" example:"fail"`
	Data   string `json:"data" example:"the error explained here"`
}
