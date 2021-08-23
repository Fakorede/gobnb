package handlers

import (
	"github.com/fakorede/gobnb/internal/config"
	"github.com/fakorede/gobnb/internal/driver"
	"github.com/fakorede/gobnb/internal/repository"
	"github.com/fakorede/gobnb/internal/repository/dbrepository"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepository.NewPostgresRepo(db.SQL, a),
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}
