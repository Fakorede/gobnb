package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCsrfTokenMiddleware(t *testing.T) {
	var myH myHandler
	h := CsrfTokenMiddleware(&myH)

	switch v := h.(type) {
	case http.Handler:
		// test passed, do nothng
	default:
		t.Error(fmt.Sprintf("Type is not http.Handler, but is type %T", v))
	}
}

func TestSessionMiddleware(t *testing.T) {
	var myH myHandler
	h := SessionMiddleware(&myH)

	switch v := h.(type) {
	case http.Handler:
		// test passed, do nothng
	default:
		t.Error(fmt.Sprintf("Type is not http.Handler, but is type %T", v))
	}
}
