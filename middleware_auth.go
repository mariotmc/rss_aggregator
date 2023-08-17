package main

import (
	"fmt"
	"net/http"

	"github.com/mariotmc/rss_aggregator/internal/auth"
	"github.com/mariotmc/rss_aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetUserByAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Couldn't get the user: %v", err))
			return
		}

		handler(w, r, user)
	}
}