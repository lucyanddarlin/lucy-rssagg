package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lucyanddarlin/rssagg/internal/databases"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user databases.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder((r.Body))

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		if params.Name == "" {
			responseWithErr(w, 400, fmt.Sprintln("Params name is missing"))
			return
		}
		if params.URL == "" {
			responseWithErr(w, 400, fmt.Sprintln("Params url is missing"))
			return
		}
		responseWithErr(w, 400, fmt.Sprintln("Error parsing JSON:", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), databases.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		responseWithErr(w, 400, fmt.Sprintln("Couldn't create feed", err))
		return
	}

	responseWithJson(w, 201, convertDatabasesFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		responseWithErr(w, 400, fmt.Sprintln("Couldn't get feeds", err))
		return
	}

	responseWithJson(w, 201, convertDatabasesFeedsToFeeds(feeds))
}
