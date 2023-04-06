package models

type Person struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone"`
}

type JsonResponse struct {
    Type    string `json:"type"`
    Data    []Person `json:"data"`
    Message string `json:"message"`
}
