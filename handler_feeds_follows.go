package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/lucyanddarlin/rssagg/internal/databases"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user databases.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder((r.Body))

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		if params.FeedID == uuid.Nil {
			responseWithErr(w, 400, fmt.Sprintln("Params feed_id is missing or error"))
			return
		}
		responseWithErr(w, 400, fmt.Sprintln("Error parsing JSON:", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), databases.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		responseWithErr(w, 400, fmt.Sprintln("Couldn't create feed_follows", err))
		return
	}

	responseWithJson(w, 201, convertDatabasesFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user databases.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		responseWithErr(w, 400, fmt.Sprintln("Couldn't create feed_follows", err))
		return
	}

	responseWithJson(w, 200, convertDatabasesFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user databases.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		responseWithErr(w, 400, fmt.Sprintln("Couldn't prase the feed follow id", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollows(r.Context(), databases.DeleteFeedFollowsParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		responseWithErr(w, 400, fmt.Sprintln("Couldn't unFollow", err))
		return
	}

	responseWithJson(w, 200, struct{}{})
}
