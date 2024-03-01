package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/wipdev-tech/blogagg/internal/database"
)

type follow struct {
	ID        uuid.UUID `json:"id"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (api *apiConfig) handleFollowsCreate(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	inFollow := struct {
		FeedID uuid.UUID `json:"feed_id"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&inFollow)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Malformed request body")
		return
	}

	dbFollow, err := api.DB.CreateFollow(r.Context(), database.CreateFollowParams{
		ID:        uuid.New(),
		FeedID:    inFollow.FeedID,
		UserID:    dbUser.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		respondWithError(w, http.StatusConflict, "Problem creating the follow: "+err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, dbFollowToFollow(dbFollow))
}

func dbFollowToFollow(f database.Follow) follow {
	return follow{
		ID:        f.ID,
		FeedID:    f.FeedID,
		UserID:    f.UserID,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
	}
}
