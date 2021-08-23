package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"general-quarters", "/general-quarters", "GET", []postData{}, http.StatusOK},
	{"major-suites", "/major-suites", "GET", []postData{}, http.StatusOK},
	{"search-availability", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"make-reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"search-availability", "/search-availability", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-02"},
	}, http.StatusOK},
	{"search-availability-json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-02"},
	}, http.StatusOK},
	{"make-reservation", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "John"},
		{key: "last_name", value: "Doe"},
		{key: "email", value: "jdoe@test.com"},
		{key: "phone", value: "333-333-3333"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)

	defer ts.Close()

	for _, v := range theTests {
		if v.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + v.url)
			if err != nil {
				t.Log(err)
				t.Fail()
			}

			if resp.StatusCode != 200 {
				t.Errorf("for %s, expected %d but got %d", v.name, v.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, p := range v.params {
				values.Add(p.key, p.value)
			}

			resp, err := ts.Client().PostForm(ts.URL+v.url, values)
			if err != nil {
				t.Log(err)
				t.Fail()
			}

			if resp.StatusCode != 200 {
				t.Errorf("for %s, expected %d but got %d", v.name, v.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
