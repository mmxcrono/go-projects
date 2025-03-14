package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/mmxcrono/goapi/api"
	"github.com/mmxcrono/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(writer http.ResponseWriter, request *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, request.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHanddler(writer)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()

	if err != nil {
		log.Error(err)
		api.InternalErrorHanddler(writer)
		return
	}

	var coinDetails *tools.CoinDetails
	coinDetails = (*database).GetUserCoinDetails(params.Username)

	if coinDetails == nil {
		log.Error("Coin details not found")
		api.InternalErrorHanddler(writer)
		return
	}

	var response = api.CoinBalanceResponse {
		Balance: (*coinDetails).Coins,
		Code: http.StatusOK,
	}

	writer.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(writer).Encode(response)

	if err != nil {
		log.Error(err)
		api.InternalErrorHanddler(writer)
		return
	}
	
}