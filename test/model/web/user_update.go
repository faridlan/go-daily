package web

type UserUpdate struct {
	Id       int    `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required, min=1,max=100"`
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required, min=8"`
}
