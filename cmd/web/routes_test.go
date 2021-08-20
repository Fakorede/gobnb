package main

import (
	"fmt"
	"testing"

	"github.com/fakorede/gobnb/internal/config"
	"github.com/go-chi/chi/v5"
)

func TestRoutes(t *testing.T) {
	var app *config.AppConfig

	mux := routes(app)

	switch v := mux.(type) {
	case *chi.Mux:
		// test passed, do nothing
	default:
		t.Error(fmt.Sprintf("Type is not *chi.Mux, but is type %T", v))
	}
}
