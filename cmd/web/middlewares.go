package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// CsrfTokenMiddleware adds CSRF protection to all post requests
func CsrfTokenMiddleware(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionMiddleware loads and saves the session on every request
func SessionMiddleware(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
