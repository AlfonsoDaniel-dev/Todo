package DTO

import (
	"github.com/google/uuid"
	"time"
	"todoApp-backend/src/models"
)

type CreateOrganization struct {
	OwnerId     uuid.UUID `json:"ownerId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type GetOrganizationPrivateData struct {
	Id            uuid.UUID                             `json:"id"`
	Owner         models.User                           `json:"owner"`
	Name          string                                `json:"name"`
	Description   string                                `json:"description"`
	Collaborators []models.User                         `json:"collaborators"`
	Projects      []models.Project[models.Organization] `json:"projects"`
	CreatedAt     time.Time                             `json:"createdAt"`
}

type GetOrganizationPublicData struct {
	Id            uuid.UUID     `json:"id"`
	Name          string        `json:"name"`
	Collaborators []models.User `json:"collaborators"`
	CreatedAt     time.Time     `json:"createdAt"`
}
