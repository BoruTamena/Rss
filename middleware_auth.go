package main

import (
	"fmt"
	"net/http"

	"github.com/BoruTamena/RssAggergetor/internal/auth"
	"github.com/BoruTamena/RssAggergetor/internal/database"
)

// creating our own custom func type signiture
type authedHader func(http.ResponseWriter, *http.Request, database.User)

// creating a middleware function

func (apicfg *apiConfig) middlewareAuth(handler authedHader) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		apikey, err := auth.GetApiKey(r.Header)

		if err != nil {
			respondWithErr(w, 400, fmt.Sprintf("auth error: %v	", err))
			return
		}

		user, err := apicfg.db.GetUserByApiKey(r.Context(), apikey)

		if err != nil {
			respondWithErr(w, 400, fmt.Sprintf("couldn't get user :%v", err))
			return
		}

		handler(w, r, user)
	}
}
