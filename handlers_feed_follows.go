package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/enesonus/go-rss-aggregator/internal/db"
	"github.com/go-chi/chi/v5"
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

func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user db.User) {
	feed_follow_id_str := chi.URLParam(r, "feed_follow_id")
	feed_follow_id, err := uuid.Parse(feed_follow_id_str)
	if err != nil {
		respondwithJSON(w, 400, map[string]string{"Couldn't parse feed_follow_id.\nerror": fmt.Sprintf("%v", err)})
		return
	}

	params := db.DeleteFeedFollowParams{
		ID:     feed_follow_id,
		UserID: user.ID,
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), params)
	if err != nil {
		respondwithJSON(w, 400, map[string]string{"Couldn't delete Feed Follow\nerror": fmt.Sprintf("%v", err)})
		return
	}
	respondwithJSON(w, 200, map[string]string{"message": fmt.Sprintf("Feed follow %v deleted.", feed_follow_id)})
}
