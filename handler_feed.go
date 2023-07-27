package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/lmirenda/RSS-aggregator/internal/database"
	"net/http"
	"time"
)

func (apiCfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Bad request: %v", err))
	}

	dbFeed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    dbUser.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt Create user: %v", err))
		return
	}

	feed := databaseFeedToFeed(dbFeed)

	respondWithJson(w, 201, feed)
}
