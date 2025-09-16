package main

import (
	"fmt"
	"net/http"
	"sample-project/database"
	hanlder "sample-project/handler"
	"sample-project/repository"
	servicev1 "sample-project/service"

	"github.com/gorilla/mux"
)

func main() {
	Router := mux.NewRouter()
	dynamoDB := database.ConnectAWS("us-east-1")
	userRepo := repository.NewUserRepository(dynamoDB)
	Service := servicev1.NewSampleService(userRepo)
	Handler := hanlder.NewSameplHandler(Service)
	Router.HandleFunc("/emp", Handler.CreateEmployee).Methods(http.MethodPost)
	Router.HandleFunc("/users", Handler.CreateUser).Methods(http.MethodPost)
	Router.HandleFunc("/users/all", Handler.GetAllUsers).Methods(http.MethodGet)
	Router.HandleFunc("/users", Handler.UpdateUser).Methods(http.MethodPut)
	Router.HandleFunc("/users/{user_id}/orders/{order_id}", Handler.DeleteUser).Methods(http.MethodDelete)
	Router.HandleFunc("/users/phone/{phone}", Handler.GetUsersByPhone).Methods(http.MethodGet)


	fmt.Println("listinning 8000")
	http.ListenAndServe(":8000", Router)

}
