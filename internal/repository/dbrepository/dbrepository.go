package dbrepository

import (
	"database/sql"

	"github.com/fakorede/gobnb/internal/config"
	"github.com/fakorede/gobnb/internal/repository"
)

type postgresRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresRepo{
		App: a,
		DB:  conn,
	}
}
