package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/BoruTamena/RssAggergetor/internal/auth"
	"github.com/BoruTamena/RssAggergetor/internal/database"
	"github.com/google/uuid"
)

func (apicf *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameter struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	parms := parameter{}
	err := decoder.Decode(&parms)

	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Error parsing JSON %v", err))
		return
	}

	user, err := apicf.db.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      parms.Name,
	})

	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("couldn't create user %v	", err))
		return
	}
	respondWithJson(w, 201, databaseUsertoUser(user))

}

func (apicf *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request) {

	apikey, err := auth.GetApiKey(r.Header)

	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("auth error: %v	", err))
		return
	}

	user, err := apicf.db.GetUserByApiKey(r.Context(), apikey)

	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("couldn't get user :%v", err))
		return
	}

	respondWithJson(w, 200, databaseUsertoUser(user))
}
