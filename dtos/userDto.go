package dtos

import "scores-backend/models"

type userDto struct {
	ID              uint          `json:"id"`
	Email           string        `json:"email"`
	Player          models.Player `json:"player"`
	PlayerID        uint          `json:"playerId"`
	ProfileImageURL string        `json:"profileImageUrl"`
}
