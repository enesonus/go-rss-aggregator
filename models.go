package main

import (
	"time"

	"github.com/enesonus/go-rss-aggregator/internal/db"
	"github.com/google/uuid"
)


type User struct {
	ID uuid.UUID `json:"id"`
	Username string `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	APIKey string `json:"api_key"`
}

func databaseUserToUser(dbUser db.User) User {
	return User{
		ID: dbUser.ID,
		Username: dbUser.Username,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		APIKey: dbUser.ApiKey,
	}
}