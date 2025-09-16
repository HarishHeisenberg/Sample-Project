package repository

import (
	"context"
	"fmt"
	"sample-project/dto"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UserRepository struct {
	db *dynamodb.Client
}
type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user dto.User) error
	GetAllUsers(ctx context.Context) ([]dto.User, error)
	UpdateUser(ctx context.Context, user dto.User) error
	DeleteUser(ctx context.Context, userID int, orderID int) error
	GetUsersByPhone(ctx context.Context, phoneNumber int) ([]dto.User, error)
}

func NewUserRepository(db *dynamodb.Client) UserRepositoryInterface {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user dto.User) error {
	av, err := attributevalue.MarshalMap(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user: %v", err)
	}

	_, err = r.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("users"),
		Item:      av,
	})
	if err != nil {
		return fmt.Errorf("failed to put item: %v", err)
	}
	return nil
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]dto.User, error) {
	var users []dto.User

	// Scan the whole table
	output, err := r.db.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String("users"),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to scan table: %v", err)
	}

	// Unmarshal DynamoDB items into Go struct
	err = attributevalue.UnmarshalListOfMaps(output.Items, &users)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal items: %v", err)
	}

	return users, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user dto.User) error {
	_, err := r.db.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String("users"),
		Key: map[string]types.AttributeValue{
			"userId":  &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", user.UserID)},
			"orderId": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", user.OrderID)},
		},
		// Update only the attributes you want to change
		UpdateExpression: aws.String("SET orderDate = :d,phoneNumber = :p"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":d": &types.AttributeValueMemberS{Value: user.OrderDate},
			":p": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", user.PhoneNumber)},
		},
		ReturnValues: types.ReturnValueUpdatedNew, // return updated values
	})
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}
	return nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, userID int, orderID int) error {
	_, err := r.db.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String("users"),
		Key: map[string]types.AttributeValue{
			"userId":  &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", userID)},
			"orderId": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", orderID)},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}

func (r *UserRepository) GetUsersByPhone(ctx context.Context, phoneNumber int) ([]dto.User, error) {
	input := &dynamodb.QueryInput{
		TableName:              aws.String("users"),
		IndexName:              aws.String("phoneNumber-orderId-index"), // GSI name
		KeyConditionExpression: aws.String("phoneNumber = :p"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":p": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", phoneNumber)},
		},
	}

	out, err := r.db.Query(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to query by phoneNumber: %v", err)
	}

	var users []dto.User
	err = attributevalue.UnmarshalListOfMaps(out.Items, &users)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal query result: %v", err)
	}

	return users, nil
}
