package hanlder

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sample-project/dto"
	servicev1 "sample-project/service"
	"strconv"

	"github.com/gorilla/mux"
)

type HandlerService struct {
	SampleService servicev1.SampleServiceInterface
}

func NewSameplHandler(sampleService servicev1.SampleServiceInterface) HandlerService {
	return HandlerService{
		SampleService: sampleService,
	}

}

func (lh HandlerService) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var Empl dto.ApiRequest
	err := json.NewDecoder(r.Body).Decode(&Empl)
	if err != nil {
		writeResponse(w, dto.ApiResponse{
			Status:  400,
			Message: err.Error()})
		return
	}
	Data := lh.SampleService.SimplePost(Empl)
	writeResponse(w, Data)

}
func (h HandlerService) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		writeResponse(w, dto.ApiResponse{
			Status:  400,
			Message: err.Error()})
		return
	}

	Data := h.SampleService.RegisterUser(context.Background(), user)
	writeResponse(w, Data)
}
func (h *HandlerService) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.SampleService.GetAllUsers(context.Background())
	if err != nil {
		writeResponse(w, err)
		return
	}
	writeResponse(w, users)
}
func (h HandlerService) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		writeResponse(w, dto.ApiResponse{
			Status:  400,
			Message: err.Error()})
		return
	}
	Data := h.SampleService.UpdateUser(context.Background(), user)
	writeResponse(w, Data)
}
func (h HandlerService) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		writeResponse(w, dto.ApiResponse{
			Status:  400,
			Message: err.Error()})
		return
	}

	orderID, err := strconv.Atoi(vars["order_id"])
	if err != nil {
		writeResponse(w, dto.ApiResponse{
			Status:  400,
			Message: err.Error()})
		return
	}

	data := h.SampleService.DeleteUser(r.Context(), userID, orderID)
	writeResponse(w, data)
}
func (h HandlerService) GetUsersByPhone(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	phoneStr := vars["phone"]

	phone, err := strconv.Atoi(phoneStr)
	if err != nil {
		writeResponse(w, dto.ApiResponse{
			Status:  400,
			Message: err.Error()})
		return
	}

	users, errUsers := h.SampleService.GetUsersByPhone(r.Context(), phone)
	if errUsers != nil {
		writeResponse(w, err)
		return
	}

	writeResponse(w, users)

}

func writeResponse(w http.ResponseWriter, data interface{}) {
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Println("Error", err.Error())
	}
}
