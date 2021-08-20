package forms

import (
	"net/http/httptest"
	"testing"
)

func Test_FormValid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("Got invalid when shoulf have been valid")
	}
}

func Test_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("Form shows valid when required fields missing")
	}
}

func Test_FormMinLength(t *testing.T) {}
