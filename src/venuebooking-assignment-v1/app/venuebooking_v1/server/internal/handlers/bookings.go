package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/venuebooking/lib/postquery"
)

func (h *Handler) HandleMyBookingsGET(w http.ResponseWriter, r *http.Request) {
	currentUser := h.currentUser(w, r)
	if currentUser == nil {
		return
	}

	bookings, err := h.rDB.GetBookingsByUser(currentUser.ID)
	if err != nil {
		fmt.Printf("[ERROR] unable to get user bookings, err : %v", err)
		renderJSON(w, http.StatusInternalServerError,
			map[string]string{"error": err500},
		)
		return
	}

	viewArgs := map[string]interface{}{
		"Header": map[string]interface{}{
			"DisplayName": currentUser.Fname,
		},
		"DisplayName": currentUser.Fname,
		"Bookings": func() (result []interface{}) {
			for _, booking := range bookings {
				stArr := strings.Split(booking.St, " ")
				etArr := strings.Split(booking.Et, " ")

				result = append(result, map[string]interface{}{
					"Id":        booking.Id,
					"Venue":     booking.VenueName,
					"Date":      stArr[0],
					"St":        stArr[1],
					"Et":        etArr[1],
					"CustName":  booking.CustName,
					"CustPhone": booking.CustPhone,
				})
			}
			return
		}(),
	}
	renderHTML(w, "bookings.html", viewArgs)
	return
}

func (h *Handler) HandleEditMyBookingGET(w http.ResponseWriter, r *http.Request) {
	currentUser := h.currentUser(w, r)
	if currentUser == nil {
		return
	}

	booking, err := h.rDB.GetBookingById(r.URL.Query().Get("b"))
	if err != nil {
		fmt.Printf("[ERROR] unable to get user booking, err : %v", err)
		renderJSON(w, http.StatusInternalServerError,
			map[string]string{"error": err500},
		)
		return
	}

	stArr := strings.Split(booking.St, " ")
	etArr := strings.Split(booking.Et, " ")

	venues, err := h.rDB.GetVenueList()
	if err != nil {
		h.internalServerError(w, r, errors.New(err500))
		return
	}

	viewArgs := map[string]interface{}{
		"Header": map[string]interface{}{
			"DisplayName": currentUser.Fname,
		},
		"DisplayName": currentUser.Fname,
		"Venues": func() (result []interface{}) {
			for _, v := range venues {
				result = append(result, map[string]interface{}{
					"Id":   v.ID,
					"Name": v.Name,
				})
			}
			return
		}(),
		"Booking": map[string]interface{}{
			"Id":        booking.Id,
			"Venue":     booking.VenueName,
			"Date":      stArr[0],
			"St":        stArr[1],
			"Et":        etArr[1],
			"CustName":  booking.CustName,
			"CustPhone": booking.CustPhone,
		},
	}
	renderHTML(w, "updateBooking.html", viewArgs)
	return
}

func (h *Handler) HandleMyBookingsPOST(w http.ResponseWriter, r *http.Request) {
	currentUser := h.currentUser(w, r)
	if currentUser == nil {
		return
	}

	bookingIDKey := "bookingId"
	venueKey := "venue"
	venueDateKey := "venueDate"
	venueStKey := "venueSt"
	venueEtKey := "venueEt"
	customerNameKey := "custName"
	customerPhoneKey := "custPhone"

	data, err := postquery.FormValues(r, []string{venueKey, venueDateKey, venueStKey, venueEtKey, customerNameKey, customerPhoneKey, bookingIDKey})
	if err != nil {
		renderJSON(w, http.StatusBadRequest,
			map[string]string{"error": err.Error()},
		)
		return
	}

	if data[venueKey] == "" || data[venueDateKey] == "" || venueStKey == "" || venueEtKey == "" || customerNameKey == "" || customerPhoneKey == "" {
		renderJSON(w, http.StatusBadRequest,
			map[string]string{"error": "All Fields are required!"},
		)
		return
	}

	var (
		st, et time.Time
	)

	stArr := strings.Split(data[venueStKey], ":")
	if len(stArr) == 3 {
		st, err = time.Parse(time.RFC3339, fmt.Sprintf("%vT%v+00:00", data[venueDateKey], data[venueStKey]))
	} else {
		st, err = time.Parse(time.RFC3339, fmt.Sprintf("%vT%v:00+00:00", data[venueDateKey], data[venueStKey]))
	}
	if err != nil {
		fmt.Printf("[ERROR] unable to parse start time, err : %v", err)
		renderJSON(w, http.StatusInternalServerError,
			map[string]string{"error": err500},
		)
		return
	}

	etArr := strings.Split(data[venueEtKey], ":")
	if len(etArr) == 3 {
		et, err = time.Parse(time.RFC3339, fmt.Sprintf("%vT%v+00:00", data[venueDateKey], data[venueEtKey]))
	} else {
		et, err = time.Parse(time.RFC3339, fmt.Sprintf("%vT%v:00+00:00", data[venueDateKey], data[venueEtKey]))
	}
	if err != nil {
		fmt.Printf("[ERROR] unable to parse end time, err : %v", err)
		renderJSON(w, http.StatusInternalServerError,
			map[string]string{"error": err500},
		)
		return
	}

	if et.Sub(st).Minutes() < 0 {
		renderJSON(w, http.StatusBadRequest,
			map[string]string{"error": "End time must be greater than Start Time!"},
		)
		return
	}

	if err = h.wDB.UpdateBooking(data[bookingIDKey], data[venueKey], st.String(), et.String(), data[customerNameKey], data[customerPhoneKey]); err != nil {
		fmt.Printf("[ERROR] unable to book a venue, err : %v", err)
		renderJSON(w, http.StatusInternalServerError,
			map[string]string{"error": err500},
		)
		return
	}

	renderJSON(w, http.StatusOK, map[string]string{"next": fmt.Sprintf("/user/bookings/edit?b=%v", data[bookingIDKey])})
}
