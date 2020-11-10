package database

import (
	"github.com/jinzhu/gorm"
)

type Sample struct {
	gorm.Model   `json:"model"`
	FileName     string `json:"fileName,omitempty"`
	Path         string `json:"path,omitempty"`
	Repository   string `json:"repository,omitempty"`
	RepositoryId int64  `json:"repositoryId,omitempty"`
	Fork         bool   `json:"fork,omitempty"`
	Content      string `json:"content:omitempty"`
}
