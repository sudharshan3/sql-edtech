package models

type Student struct {
	Id       int    `json:"student_id"`
	Name     string `json:"student_name"`
	Email    string `json:"student_email"`
	Phone    string `json:"student_phone"`
	Password string `json:"student_pass,omitempty"`
}

type Response struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Student `json:"data,omitempty"`
}

type LoginInfo struct {
	Email string `json:"student_email"`
	Token string `json:"token"`
}
type LoginResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    LoginInfo `json:"data,omitempty"`
}
