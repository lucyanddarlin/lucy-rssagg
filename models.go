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

type Feed struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

func convertDatabasesFeedToFeed(feed databases.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
	}
}

func convertDatabasesFeedsToFeeds(dbFeeds []databases.Feed) []Feed {
	feeds := []Feed{}

	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, convertDatabasesFeedToFeed(dbFeed))
	}

	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func convertDatabasesFeedFollowToFeedFollow(feedFollow databases.FeedsFollow) FeedFollow {
	return FeedFollow{
		ID:        feedFollow.ID,
		UserID:    feedFollow.ID,
		FeedID:    feedFollow.FeedID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
	}
}

func convertDatabasesFeedFollowsToFeedFollows(dbFeedFollows []databases.FeedsFollow) []FeedFollow {
	feedFollows := []FeedFollow{}

	for _, dbFeedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, convertDatabasesFeedFollowToFeedFollow(dbFeedFollow))
	}

	return feedFollows
}
