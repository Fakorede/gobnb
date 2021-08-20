package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_FormValid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("Got invalid when should have been valid")
	}
}

func Test_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("Shows form valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("Shows form does not have required fields when it does")
	}
}

func Test_FormHas(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	has := form.Has("whatever")
	if has {
		t.Error("Shows form has field when it does not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("Shows form does not have field when it should")
	}

}

func Test_FormMinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("Form Shows minlength for non existent fields")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("Should have an error but did not get one")
	}

	postedData := url.Values{}
	postedData.Add("some_field", "some value")
	form = New(postedData)

	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("Form Shows minlength for passed length when data is shorter")
	}
	

	postedData = url.Values{}
	postedData.Add("another_field", "abc123")
	form = New(postedData)

	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("Form Shows minlength is not met length when it is")
	}

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("Should not have an error but got one")
	}
}

func Test_FormIsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("Form Shows is valid for non-existent field")
	}

	postedData = url.Values{}
	postedData.Add("email", "jdoe@test.com")
	form = New(postedData)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("Form Shows invalid email when email is valid")
	}

	postedData = url.Values{}
	postedData.Add("email", "x")
	form = New(postedData)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("Form Shows valid email when email is invalid")
	}
}
