package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/fakorede/gobnb/internal/forms"
	"github.com/fakorede/gobnb/internal/helpers"
	"github.com/fakorede/gobnb/internal/models"
	"github.com/fakorede/gobnb/internal/render"
)

// Reservation renders the form for making reservations
func (rh *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	res, ok := rh.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(w, errors.New("cannot get reservation from session"))
		return
	}

	room, err := rh.DB.GetRoomByID(res.RoomID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	res.Room.RoomName = room.RoomName

	rh.App.Session.Put(r.Context(), "reservation", res)

	start_date := res.StartDate.Format("2006-01-02")
	end_date := res.EndDate.Format("2006-01-02")

	stringMap := make(map[string]string)
	stringMap["start_date"] = start_date
	stringMap["end_date"] = end_date

	data := make(map[string]interface{})
	data["reservation"] = res

	render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
		StringMap: stringMap,
	})
}

// MakeReservation handles creation of reservation
func (rh *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	reservation, ok := rh.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(w, errors.New("can't get reservation from session"))
		return
	}
	
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// start_date := r.Form.Get("start_date")
	// end_date := r.Form.Get("end_date")

	// // 2021-08-23 (reservation format) | 01/02 03:04:05PM '06 -0700 (go reference time format)

	// layout := "2006-01-02"
	// startDate, err := time.Parse(layout, start_date)
	// if err != nil {
	// 	helpers.ServerError(w, err)
	// 	return
	// }

	// endDate, err := time.Parse(layout, end_date)
	// if err != nil {
	// 	helpers.ServerError(w, err)
	// 	return
	// }

	// roomID, err := strconv.Atoi(r.Form.Get("room_id"))
	// if err != nil {
	// 	helpers.ServerError(w, err)
	// 	return
	// }

	reservation.FirstName = r.Form.Get("first_name")
	reservation.LastName = r.Form.Get("last_name")
	reservation.Email = r.Form.Get("email")
	reservation.Phone = r.Form.Get("phone")

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3)
	form.MinLength("last_name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
	}

	newReservationID, err := rh.DB.InsertReservation(reservation)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	restriction := models.RoomRestriction{
		StartDate:     reservation.StartDate,
		EndDate:       reservation.EndDate,
		RoomID:        reservation.RoomID,
		ReservationID: newReservationID,
		RestrictionID: 1,
	}

	err = rh.DB.InsertRoomRestriction(restriction)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	rh.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func (rh *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := rh.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		rh.App.ErrorLog.Println("Cannot get item from session")
		rh.App.Session.Put(r.Context(), "error", "Cannot get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	rh.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation

	start_date := reservation.StartDate.Format("2006-01-02")
	end_date := reservation.EndDate.Format("2006-01-02")

	stringMap := make(map[string]string)

	stringMap["start_date"] = start_date
	stringMap["end_date"] = end_date


	render.Template(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
		StringMap: stringMap,
	})
}

func (rh *Repository) ChooseRoom(w http.ResponseWriter, r *http.Request) {
	roomID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	res, ok := rh.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(w, err)
		return
	}

	res.RoomID = roomID

	rh.App.Session.Put(r.Context(), "reservation", res)

	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)
}
