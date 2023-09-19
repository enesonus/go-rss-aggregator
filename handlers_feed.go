package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/enesonus/go-rss-aggregator/internal/db"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user db.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	params := parameters{}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondwithJSON(w, 400, map[string]string{"error": fmt.Sprintf("%v", err)})
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), db.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		respondwithJSON(w, 400, map[string]string{"error": fmt.Sprintf("%v", err)})
		return
	}

	respondwithJSON(w, 201, databaseFeedToFeed(feed))
}
