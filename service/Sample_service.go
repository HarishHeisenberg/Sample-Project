package services

import (
	"fmt"
	"sample-project/dto"
)

type SampleService struct{}

type SampleServiceInterface interface {
	SimplePost(Emplyoee dto.ApiRequest) dto.ApiResponse
}

func NewSampleService() SampleServiceInterface {
	return SampleService{}
}

func (lh SampleService) SimplePost(Emplyoee dto.ApiRequest) dto.ApiResponse {
	fmt.Println(Emplyoee)
	return dto.ApiResponse{
		Status:  200,
		Message: "success",
	}

}
