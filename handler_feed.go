package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/BoruTamena/RssAggergetor/internal/database"
	"github.com/google/uuid"
)

func (apicf *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameter struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	parms := parameter{}
	err := decoder.Decode(&parms)

	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Error parsing JSON %v", err))
		return
	}

	feed, err := apicf.db.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      parms.Name,
		Url:       parms.URL,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("couldn't create user %v	", err))
		return
	}
	respondWithJson(w, 201, feed)

}
