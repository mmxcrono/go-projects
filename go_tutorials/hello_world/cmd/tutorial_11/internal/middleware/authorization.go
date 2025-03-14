package middleware

import (
	"errors"
	"net/http"

	"github.com/mmxcrono/goapi/api"
	"github.com/mmxcrono/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var errUnauthorizezd = errors.New("invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		var username string = request.URL.Query().Get("username")
		var token string = request.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(errUnauthorizezd)
			api.RequestErrorHandler(responseWriter, errUnauthorizezd)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()

		if err != nil {
			api.InternalErrorHanddler(responseWriter)
			return
		}

		var loginDetails *tools.LoginDetails = (*database).GetUserLoginDetails(username)

		if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
			log.Error(errUnauthorizezd)
			api.RequestErrorHandler(responseWriter, errUnauthorizezd)
			return
		}
		next.ServeHTTP(responseWriter, request)
	})
}