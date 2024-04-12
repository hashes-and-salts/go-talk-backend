package models

type User struct {
	Id        uint   `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"userEmail"`
	Password  string `json:"userPassword"`
}
