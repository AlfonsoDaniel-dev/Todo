package DTO

import "github.com/google/uuid"

type CreateProjectDTO struct {
	CreatorId     uuid.UUID
	Name          string
	Description   string
	Collaborators []uuid.UUID
}

type UpdateProjectNameDTO struct {
}
