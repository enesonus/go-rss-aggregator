package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/enesonus/go-rss-aggregator/internal/db"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request){
	type parameters struct {
		Name string `json:"username"`
		URL string `json:"url"`
		UserID string `json:"user_id"`
	}

	params := parameters{}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondwithJSON(w, 400, map[string]string{"error": fmt.Sprintf("%v", err)})
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), db.CreateFeedParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name: params.Name,
			Url: params.URL, 
			UserID: uuid.MustParse(params.UserID),
		})
		
	if err != nil {
		respondwithJSON(w, 400, map[string]string{"error": fmt.Sprintf("%v", err)})
	}

	respondwithJSON(w, 201, databaseFeedToFeed(feed))
}