package main

type ClientProfile struct {
	Email string
	Id    string
	Name  string
	Token string
}

var database = map[string]ClientProfile{
	"user1": {
		Email: "test@email.com",
		Id:    "user1",
		Name:  "Test User",
		Token: "123",
	},
}