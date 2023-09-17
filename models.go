package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/lucyanddarlin/rssagg/internal/databases"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	APIKey    string    `json:"api_key"`
}

func convertDatabasesUserToUser(user databases.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		APIKey:    user.ApiKey,
	}

}
