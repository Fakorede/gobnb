package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/fakorede/gobnb/internal/helpers"
	"github.com/fakorede/gobnb/internal/models"
	"github.com/fakorede/gobnb/internal/render"
)

// Availability renders the search availability page
func (rh *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// CheckAvailability checks for available room btw date range
func (rh *Repository) CheckAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	// 2021-08-23 (reservation format) | 01/02 03:04:05PM '06 -0700 (go reference time format)
	layout := "2006-01-02"
	startDate, err := time.Parse(layout, start)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	endDate, err := time.Parse(layout, end)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	rooms, err := rh.DB.SearchAvailabilityForAllRooms(startDate, endDate) 
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	if len(rooms) == 0 {
		// rh.App.InfoLog.Println("No Available Rooms")
		rh.App.Session.Put(r.Context(), "error", "No Available Rooms")
		http.Redirect(w, r, "/search-availability", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["rooms"] = rooms

	res := models.Reservation{
		StartDate: startDate,
		EndDate: endDate,
	}

	rh.App.Session.Put(r.Context(), "reservation", res)

	render.Template(w, r, "available-rooms.page.tmpl", &models.TemplateData{
		Data: data,
	})

	// w.Write([]byte(fmt.Sprintf("Start date is %s and End date is %s", start, end)))
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
		helpers.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
