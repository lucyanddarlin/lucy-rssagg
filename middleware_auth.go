package main

import (
	"fmt"
	"net/http"

	"github.com/lucyanddarlin/rssagg/internal/auth"
	"github.com/lucyanddarlin/rssagg/internal/databases"
)

type authHandler func(w http.ResponseWriter, r *http.Request, user databases.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responseWithErr(w, 403, fmt.Sprintln("Authorization Error:", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			responseWithErr(w, 400, fmt.Sprintln("Couldn't get user", err))
			return
		}
		handler(w, r, user)
	}
}
