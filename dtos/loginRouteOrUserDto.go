package dtos

import "scores-backend/models"

type LoginRouteOrUserDto struct {
	LoginRoute string       `json:"loginRoute"`
	User       *models.User `json:"user"`
}
