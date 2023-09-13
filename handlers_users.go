package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/enesonus/go-rss-aggregator/internal/auth"
	"github.com/enesonus/go-rss-aggregator/internal/db"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request){
	type parameters struct {
		Username string `json:"username"`
	}

	params := parameters{}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondwithJSON(w, 400, map[string]string{"error": fmt.Sprintf("%v", err)})
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), db.CreateUserParams{
			ID: uuid.New(),
			Username: params.Username,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
		
	if err != nil {
		respondwithJSON(w, 400, map[string]string{"error": fmt.Sprintf("%v", err)})
	}

	respondwithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUserByKey(w http.ResponseWriter, r *http.Request){
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondwithJSON(w, 403, map[string]string{"error": fmt.Sprintf("%v", err)})
		return
	}
	user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey) 
	if err != nil {
		respondwithJSON(w, 400, map[string]string{"error": fmt.Sprintf("%v", err)})
		return
	}
	respondwithJSON(w, 200, databaseUserToUser(user))
}