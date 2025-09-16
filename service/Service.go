package servicev1

import (
	"context"
	"fmt"
	"sample-project/dto"
	"sample-project/repository"
)

type SampleService struct {
	repo repository.UserRepositoryInterface
}

type SampleServiceInterface interface {
	SimplePost(Emplyoee dto.ApiRequest) dto.ApiResponse
	RegisterUser(ctx context.Context, user dto.User) dto.ApiResponse
	GetAllUsers(ctx context.Context) (*[]dto.User, *dto.ApiResponse)
	UpdateUser(ctx context.Context, user dto.User) dto.ApiResponse
	DeleteUser(ctx context.Context, userID int, orderID int) dto.ApiResponse
	GetUsersByPhone(ctx context.Context, phoneNumber int) (*[]dto.User, *dto.ApiResponse)
}

func NewSampleService(repo repository.UserRepositoryInterface) SampleServiceInterface {
	return &SampleService{repo: repo}
}

func (lh SampleService) SimplePost(Emplyoee dto.ApiRequest) dto.ApiResponse {
	fmt.Println(Emplyoee)
	return dto.ApiResponse{
		Status:  200,
		Message: "success",
	}

}

func (s *SampleService) RegisterUser(ctx context.Context, user dto.User) dto.ApiResponse {
	if user.UserID == 0 {
		return dto.ApiResponse{
			Status:  400,
			Message: "userId is required",
		}
	}
	err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return dto.ApiResponse{
			Status:  500,
			Message: err.Error(),
		}
	}
	return dto.ApiResponse{
		Status:  200,
		Message: "successfully created",
	}
}

func (s *SampleService) GetAllUsers(ctx context.Context) (*[]dto.User, *dto.ApiResponse) {
	data, err := s.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, &dto.ApiResponse{
			Status:  500,
			Message: err.Error(),
		}
	}
	return &data, nil

}
func (s *SampleService) UpdateUser(ctx context.Context, user dto.User) dto.ApiResponse {
	if user.UserID == 0 {
		return dto.ApiResponse{
			Status:  400,
			Message: "userId is required",
		}
	}
	err := s.repo.UpdateUser(ctx, user)
	if err != nil {
		return dto.ApiResponse{
			Status:  500,
			Message: err.Error(),
		}
	}
	return dto.ApiResponse{
		Status:  200,
		Message: "successfully Updated",
	}
}

func (s *SampleService) DeleteUser(ctx context.Context, userID int, orderID int) dto.ApiResponse {
	err := s.repo.DeleteUser(ctx, userID, orderID)
	if err != nil {
		return dto.ApiResponse{
			Status:  500,
			Message: err.Error(),
		}
	}
	return dto.ApiResponse{
		Status:  200,
		Message: "successfully Updated",
	}
}

func (s *SampleService) GetUsersByPhone(ctx context.Context, phoneNumber int) (*[]dto.User, *dto.ApiResponse) {
	data, err := s.repo.GetUsersByPhone(ctx, phoneNumber)
	if err != nil {
		return nil, &dto.ApiResponse{
			Status:  500,
			Message: err.Error(),
		}
	}
	return &data, nil
}
