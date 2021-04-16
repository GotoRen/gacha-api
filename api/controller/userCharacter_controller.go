package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fumist23/game-api/database"
)

func GetUserCharacters(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	token := r.Header.Get("x-token")

	isValidToken := database.VerifyToken(ctx, token)
	if !isValidToken {
		log.Printf("this token is invalid: %v", token)
		w.WriteHeader(http.StatusUnauthorized)
	}

	user, err := database.GetUser(ctx, token)
	if err != nil {
		log.Printf("failed to get user: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// userCharacters, err := database.GetUserCharactersByID(ctx, user.Id)
	userCharacters, err := database.GetUserCountedCharactersByID(ctx, user.Id)
	if err != nil {
		log.Printf("failed to get userCharacters: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(userCharacters); err != nil {
		log.Printf("failed to encode userCharacters")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
