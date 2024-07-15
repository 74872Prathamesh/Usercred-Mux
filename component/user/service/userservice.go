package service

import (
	"fmt"
	"usercred/model"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) GetUsers() ([]*model.User, error) {
	//Validate
	//logic here communicate with model
	//return data
	fetchedUsers, err := model.GetUsers()
	if err != nil {
		return nil, err
	}
	fmt.Println(fetchedUsers, " all users in service ")
	return fetchedUsers, nil

}

func (u *UserService) CreateUser(user *model.User) (bool, error) {
	//Validate
	//logic here communicate with model
	//return data
	var allUser []*model.User
	err := user.ValidateUser()
	if err != nil {
		return false, err
	}

	addedUser, err := user.CreateUser()
	if err != nil {
		return false, err
	}

	allUser = append(allUser, addedUser)
	fmt.Println(allUser, " all user in service method !!")
	return true, nil

}

func (u *UserService) UpdateUser(userID string, user *model.User) (*model.User, error) {
	//Validate
	//logic here communicate with model
	//return data
	err := user.ValidateUser()
	if err != nil {
		return nil, err
	}
	updatedUser, err := model.UpdateUser(userID, user)
	if err != nil {
		return nil, err
	}
	fmt.Println(updatedUser, " updated users in service ")
	return updatedUser, nil

}

func (u *UserService) DeleteUser(userID string) (*model.User, error) {
	//Validate
	//logic here communicate with model
	//return data
	deletedUser, err := model.DeleteUser(userID)
	if err != nil {
		return nil, err
	}
	fmt.Println(deletedUser, " deleted user in service ")
	return deletedUser, nil

}
