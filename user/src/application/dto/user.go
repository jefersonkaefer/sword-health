package dto

type UserCreateDTO struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	RePassword string `json:"re-password"`
	Role       string `json:"role"`
}

func CreateUser(
	firstName string,
	lastName string,
	email string,
	role string,
	password string,
	rePassword string,
) UserCreateDTO {
	return UserCreateDTO{
		FirstName:  firstName,
		LastName:   lastName,
		Email:      email,
		Role:       role,
		Password:   password,
		RePassword: rePassword,
	}
}

type FindUser struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	Role      string
}
