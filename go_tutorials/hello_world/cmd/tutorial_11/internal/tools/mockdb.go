package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map [string] LoginDetails {
	"testuser": {
		Username: "testuser",
		AuthToken: "testtoken",
	},
}

var mockCoinDetails = map[string] CoinDetails {
	"testuser": {
		Coins: 1000,
		Username: "testuser",
	},
}

func (db *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second)

	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDB) GetUserCoinDetails(username string) *CoinDetails {
	time.Sleep(time.Second)

	clientData, ok := mockCoinDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDB) SetupDatabase() error {
	return nil
}