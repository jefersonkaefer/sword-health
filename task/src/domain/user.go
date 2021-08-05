package domain

type User struct {
	id   int
	role string
}

func (User) Load(id int, role string) *User {
	return &User{
		id:   id,
		role: role,
	}
}

func (u *User) IsManager() bool {
	return u.role == "manager"
}

func (u *User) GetId() int {
	return u.id
}
