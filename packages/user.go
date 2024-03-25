package user

import "fmt"

type user struct {
	name string
	age  int
}

// NewUser is a constructor of user
func NewUser(name string, age int) *user {
	return &user{
		name,
		age,
	}
}

// PrintInfo is a method of user
func (user *user) PrintInfo() {
	fmt.Println(user.name, user.age)
}

// SetName is a method of user
func (u *user) SetName(name string) {
	u.name = name
}