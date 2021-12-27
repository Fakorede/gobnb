package handlers

import (
	"github.com/fakorede/gobnb/internal/config"
	"github.com/fakorede/gobnb/internal/driver"
	"github.com/fakorede/gobnb/internal/repository"
	"github.com/fakorede/gobnb/internal/repository/dbrepository"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type - we'll be needing the app config and db in our handlers
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo creates a new repository - instantiates the repo to be used by the handler with the app config and db driver
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepository.NewPostgresRepo(db.SQL, a),
	}
}

// NewHandlers sets the repository for the handlers - initializes the handler with the stuffs it needs
func NewHandlers(r *Repository) {
	Repo = r
}
