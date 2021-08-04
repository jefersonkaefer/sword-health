package domain

import (
	"log"
	"sword-health/users/application/data_model"

	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	id        uint
	firstName string
	lastName  string
	email     string
	password  string
	fullName  string
	role      string
	isDeleted bool
}

func Create(
	firstName string,
	lastName string,
	email string,
	password string,
	role string,
) UserModel {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		log.Panic("Error ", err.Error())
	}

	u := UserModel{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		password:  string(hash),
		role:      role,
	}

	u.isDeleted = false
	return u
}

func (u *UserModel) GetId() uint {
	return u.id
}

func (u *UserModel) GetFirstName() string {
	return u.firstName
}

func (u *UserModel) GetLastName() string {
	return u.lastName
}

func (u *UserModel) GetEmail() string {
	return u.email
}

func (u *UserModel) GetPassword() string {
	return u.password
}

func (u *UserModel) GetRole() string {
	return u.role
}

func (u *UserModel) GetDataModel() *data_model.User {
	return &data_model.User{
		FirstName: u.firstName,
		LastName:  u.lastName,
		Email:     u.email,
		Password:  u.password,
		Role:      u.role,
	}
}

func (UserModel) Load(user *data_model.User) *UserModel {
	return &UserModel{
		id:        user.ID,
		firstName: user.FirstName,
		lastName:  user.LastName,
		password:  user.Password,
		email:     user.Email,
		role:      user.Role,
	}
}

func (u *UserModel) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
	return err == nil
}

func (u *UserModel) CheckEmail(email string) bool {
	return u.email == email
}
