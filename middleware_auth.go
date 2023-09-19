package main

import (
	"fmt"
	"net/http"

	"github.com/enesonus/go-rss-aggregator/internal/auth"
	"github.com/enesonus/go-rss-aggregator/internal/db"
)

type authedHandler func(http.ResponseWriter, *http.Request, db.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondwithJSON(w, 403, map[string]string{"error": fmt.Sprintf("%v", err)})
			return
		}
		user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondwithJSON(w, 400, map[string]string{"error": fmt.Sprintf("%v", err)})
			return
		}
		handler(w, r, user)
	}
}
