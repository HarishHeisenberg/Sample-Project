package hanlder

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sample-project/dto"
	services "sample-project/service"
)

type HandlerService struct {
	SampleService services.SampleServiceInterface
}

func NewSameplHandler(sampleService services.SampleServiceInterface) HandlerService {
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

func writeResponse(w http.ResponseWriter, data interface{}) {
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Println("Error", err.Error())
	}
}
