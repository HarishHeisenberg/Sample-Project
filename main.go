package main

import (
	"fmt"
	"net/http"
	hanlder "sample-project/handler"
	services "sample-project/service"

	"github.com/gorilla/mux"
)

func main() {
	Router := mux.NewRouter()
	Service := services.NewSampleService()
	Handler := hanlder.NewSameplHandler(Service)
	Router.HandleFunc("/emp", Handler.CreateEmployee).Methods(http.MethodPost)
	fmt.Println("listinning 8000")
	http.ListenAndServe(":8000", Router)

}
