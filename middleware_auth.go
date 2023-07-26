package main

import (
	"fmt"
	"github.com/lmirenda/RSS-aggregator/internal/auth"
	"github.com/lmirenda/RSS-aggregator/internal/database"
	"net/http"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKeyValue, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 401, fmt.Sprintf("Error in headers: %v", err))
			return
		}
		dbUser, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKeyValue)

		if err != nil {
			respondWithError(w, 401, fmt.Sprintf("Invalid API Key: %v, Error: %v", apiKeyValue, err))
			return
		}
		handler(w, r, dbUser)
	}
}
