package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fakorede/gobnb/internal/models"
	"github.com/fakorede/gobnb/internal/render"
)

func (rh *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

func (rh *Repository) CheckAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("Start date is %s and End date is %s", start, end)))
}

// jsonResponse
type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// CheckAvailabilityJSON handles request for availability and sends JSON response
func (rh *Repository) CheckAvailabilityJSON(w http.ResponseWriter, r *http.Request) {

	// start := r.Form.Get("start")
	// end := r.Form.Get("end")

	resp := jsonResponse{
		OK:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "	")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
