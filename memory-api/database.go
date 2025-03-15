package main

type ClientProfile struct {
	Id    string
	Email string
	Name  string
	Token string
}

var database = map[string]ClientProfile{
	"user1": {
		Email: "test1@email.com",
		Id:    "user1",
		Name:  "Test User 1",
		Token: "123",
	},
	"user2": {
		Email: "test2@email.com",
		Id:    "user2",
		Name:  "Test User 2",
		Token: "1234",
	},
}