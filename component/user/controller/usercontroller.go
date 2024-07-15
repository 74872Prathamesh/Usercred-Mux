package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"usercred/component/user/service"
	"usercred/model"

	"github.com/gorilla/mux"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {

	us := service.NewUserService()
	return &UserController{
		userService: us,
	}
}

func (controller *UserController) RegisterRoutes(router *mux.Router) {
	userRouter := router.PathPrefix("/user").Subrouter()

	//handle func
	userRouter.HandleFunc("/", controller.get).Methods("GET")
	userRouter.HandleFunc("/", controller.create).Methods("POST")
	userRouter.HandleFunc("/{id}", controller.update).Methods("PUT")
	userRouter.HandleFunc("/{id}", controller.delete).Methods("DELETE")

}

func (controller *UserController) get(w http.ResponseWriter, r *http.Request) {

	defer func() {
		details := recover()
		if details != nil {
			fmt.Println(details)
		}
	}()
	var allUsers []*model.User

	users, err := controller.userService.GetUsers()
	if err != nil {
		panic(err)
	}
	allUsers = users
	fmt.Println(allUsers, " all users getting in controller")
	json.NewEncoder(w).Encode(allUsers)
	w.WriteHeader(200)

}

func (controller *UserController) create(w http.ResponseWriter, r *http.Request) {

	defer func() {
		details := recover()
		if details != nil {
			fmt.Println(details)
		}
	}()
	var user *model.User
	json.NewDecoder(r.Body).Decode(&user)
	status, err := controller.userService.CreateUser(user)
	if !status {
		panic(err)
	}
	fmt.Println(user, " user added in controller method")
	json.NewEncoder(w).Encode(user)
	w.WriteHeader(200)

}

func (controller *UserController) update(w http.ResponseWriter, r *http.Request) {

	defer func() {
		details := recover()
		if details != nil {
			fmt.Println(details)
		}
	}()
	var user *model.User
	json.NewDecoder(r.Body).Decode(&user)
	params := mux.Vars(r)
	userID := (params[("id")])

	updatedUser, err := controller.userService.UpdateUser(userID, user)
	if err != nil {
		panic(err)
	}
	fmt.Println(updatedUser, " user updated in controller method")
	json.NewEncoder(w).Encode(updatedUser)
	w.WriteHeader(200)

}

func (controller *UserController) delete(w http.ResponseWriter, r *http.Request) {

	defer func() {
		details := recover()
		if details != nil {
			fmt.Println(details)
		}
	}()
	params := mux.Vars(r)
	userID := (params[("id")])

	deletedUser, err := controller.userService.DeleteUser(userID)
	if err != nil {
		panic(err)
	}
	fmt.Println(deletedUser, " user deleted in controller method")
	json.NewEncoder(w).Encode(deletedUser)
	w.WriteHeader(200)

}
