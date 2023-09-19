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

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string   `json:"name"`
	Url       string  `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
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

func databaseFeedToFeed(dbFeed db.Feed) Feed {
	return Feed{
		ID: dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name: dbFeed.Name,
		Url: dbFeed.Url,
		UserID: dbFeed.UserID,
	}
}