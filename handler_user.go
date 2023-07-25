package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/lmirenda/RSS-aggregator/internal/database"
	"net/http"
	"time"
)

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	dbUser, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	user := databaseUserToUser(dbUser)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt Create user: %v", err))
		return
	}
	respondWithJson(w, 201, user)
}
