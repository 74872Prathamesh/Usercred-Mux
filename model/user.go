package model

import (
	"errors"
	"fmt"
)

type User struct {
	Id       string
	UserName string
	Password string
}

var allUsers []*User

func (u *User) ValidateUser() error {

	defer func() {
		details := recover()
		if details != nil {
			fmt.Println(details)
		}
	}()
	if u.Id == " " || u.UserName == " " || u.Password == " " {
		panic(errors.New("invalid credentials"))
	}
	return nil

}

func (u *User) CreateUser() (*User, error) {

	defer func() {
		details := recover()
		if details != nil {
			fmt.Println(details)
		}
	}()
	if u == nil {
		panic(errors.New("invalid user"))
	}
	user := &User{
		Id:       u.Id,
		UserName: u.UserName,
		Password: u.Password,
	}
	allUsers = append(allUsers, user)
	return user, nil

}

func GetUsers() ([]*User, error) {

	if allUsers == nil {
		return nil, errors.New("could not find users ")
	}

	fmt.Println(allUsers, " allusers in model")
	return allUsers, nil

}

func UpdateUser(userID string, user *User) (*User, error) {

	if allUsers == nil {
		return nil, errors.New("could not find users ")
	}

	for _, u := range allUsers {
		if u.Id == userID {
			u.UserName = user.UserName
			u.Password = user.Password
		}
	}

	fmt.Println(user, " allusers in model")
	return user, nil

}

func DeleteUser(userID string) (*User, error) {

	if allUsers == nil {
		return nil, errors.New("could not find users ")
	}
	var deletedUser *User
	for i, u := range allUsers {
		if u.Id == userID {
			deletedUser = u
			allUsers = append(allUsers[:i], allUsers[i+1:]...)
		}

	}

	fmt.Println(deletedUser, " allusers in model")
	return deletedUser, nil

}
