package models

import (
	"github.com/google/uuid"
)

type Service struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description,omitempty"`
	VersionCount int       `json:"versionCount" db:"version_count"`
}

type SingleService struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Versions    []Version `json:"versions,omitempty"`
}
