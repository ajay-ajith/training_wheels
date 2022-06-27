package main

import "testing"

type User struct {
	Username string `json:"username"`
	Password string
	Email    string `json:"email"`
}

var registeredUsers = []User{
	{Username: "ajay", Password: "passforajay", Email: "ajay@example.com"},
	{Username: "akshay", Password: "passforakshay", Email: "akshay@example.com"},
	{Username: "rohith", Password: "passforrohith", Email: "rohith@example.com"},
}

func Authenticate(username, password string) bool {
	for _, user := range registeredUsers {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}

func TestUserAuth(t *testing.T) {
	var username string = "ajay"
	var password string = "passforajay"
	isUser := Authenticate(username, password)
	if isUser == false {
		t.Fatalf("Got unexpected result. Expected %v, got %v", true, isUser)
	}
}
