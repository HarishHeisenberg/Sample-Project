package dto

type ApiRequest struct {
	Name string `json:"name"`
}
type User struct {
	UserID      int64  `json:"userId" dynamodbav:"userId"`
	OrderID     int64  `json:"orderId" dynamodbav:"orderId"`
	OrderDate   string `json:"orderDate" dynamodbav:"orderDate"`
	PhoneNumber int    `json:"phoneNumber" dynamodbav:"phoneNumber"`
}
