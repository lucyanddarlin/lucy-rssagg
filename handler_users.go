package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lucyanddarlin/rssagg/internal/databases"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder((r.Body))

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		if params.Name == "" {
			responseWithErr(w, 400, fmt.Sprintln("Params Name is missing"))
			return
		}
		responseWithErr(w, 400, fmt.Sprintln("Error parsing JSON:", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), databases.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		responseWithErr(w, 400, fmt.Sprintln("Error when create User:", err))
		return
	}

	responseWithJson(w, 201, convertDatabasesUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user databases.User) {
	responseWithJson(w, 200, convertDatabasesUserToUser(user))
}
