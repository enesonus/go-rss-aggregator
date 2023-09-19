package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/enesonus/go-rss-aggregator/internal/db"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	params := parameters{}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondwithJSON(w, 400, map[string]string{"error": fmt.Sprintf("%v", err)})
		return
	}

	feed_follow, err := apiCfg.DB.CreateFeedFollow(r.Context(), db.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		respondwithJSON(w, 400, map[string]string{"error": fmt.Sprintf("%v", err)})
		return
	}

	respondwithJSON(w, 201, databaseFeedFollowToFeedFollow(feed_follow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request) {

	feed_follows, err := apiCfg.DB.GetFeedFollows(r.Context())
	if err != nil {
		respondwithJSON(w, 400, map[string]string{"error": fmt.Sprintf("%v", err)})
		return
	}
	respondwithJSON(w, 200, databaseFeedFollowsToFeedFollows(feed_follows))
}
