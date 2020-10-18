package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/venuebooking/lib/postquery"
)

func (h *Handler) HandleVenueListGET(w http.ResponseWriter, r *http.Request) {
	currentUser := h.currentUser(w, r)
	if currentUser == nil {
		return
	}

	venues, err := h.rDB.GetVenueList()
	if err != nil {
		fmt.Printf("error : %v", err)
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
					"Id":    v.ID,
					"Name":  v.Name,
					"Image": v.Image,
				})
			}
			return
		}(),
	}
	renderHTML(w, "venueList.html", viewArgs)
	return
}

func (h *Handler) HandleVenueDetailsGET(w http.ResponseWriter, r *http.Request) {
	currentUser := h.currentUser(w, r)
	if currentUser == nil {
		return
	}

	venueName := r.URL.Query().Get("v")
	if venueName == "" {
		renderJSON(w, http.StatusBadRequest,
			map[string]string{"error": "Missing required query parameters!"},
		)
		return
	}

	venueImage := r.URL.Query().Get("i")
	if venueImage == "" {
		renderJSON(w, http.StatusBadRequest,
			map[string]string{"error": "Missing required query parameters!"},
		)
		return
	}

	vList, err := h.rDB.GetBookingsByVenue(venueName)
	if err != nil {
		fmt.Printf("[ERROR] unable to get booked venues, err : %v", err)
		renderJSON(w, http.StatusInternalServerError,
			map[string]string{"error": err500},
		)
		return
	}

	// viewArgs handler takes in data from db/venue.go
	viewArgs := map[string]interface{}{
		"Header": map[string]interface{}{
			"DisplayName": currentUser.Fname,
		},
		"DisplayName": currentUser.Fname,
		"Venue":       venueName,
		"VenueImg":    venueImage,
		"BookedSlots": func() (result []interface{}) {
			for _, v := range vList {
				stArr := strings.Split(v.St, " ")
				etArr := strings.Split(v.Et, " ")

				result = append(result, map[string]interface{}{
					"St": fmt.Sprintf("%v    %v", stArr[0], stArr[1]),
					"Et": fmt.Sprintf("%v    %v", etArr[0], etArr[1]),
				})
			}
			return
		}(),
	}
	renderHTML(w, "venue_details.html", viewArgs)
	return
}

func (h *Handler) HandleSearchVenueGET(w http.ResponseWriter, r *http.Request) {
	currentUser := h.currentUser(w, r)
	if currentUser == nil {
		return
	}

	venues, err := h.rDB.GetSearchedVenue(r.URL.Query().Get("ven"))
	if err != nil {
		h.internalServerError(w, r, errors.New(err500))
	}

	viewArgs := map[string]interface{}{
		"Header": map[string]interface{}{
			"DisplayName": currentUser.Fname,
		},
		"DisplayName": currentUser.Fname,
		"Venues": func() (result []interface{}) {
			for _, v := range venues {
				result = append(result, map[string]interface{}{
					"Id":    v.ID,
					"Name":  v.Name,
					"Image": v.Image,
				})
			}
			return
		}(),
	}
	renderHTML(w, "searchVenue.html", viewArgs)
	return
}

func (h *Handler) HandleNewVenueBookingGET(w http.ResponseWriter, r *http.Request) {
	currentUser := h.currentUser(w, r)
	if currentUser == nil {
		return
	}
	venues, err := h.rDB.GetVenueList()
	if err != nil {
		h.internalServerError(w, r, errors.New(err500))
	}

	viewArgs := map[string]interface{}{
		"Header": map[string]interface{}{
			"DisplayName": currentUser.Fname,
		},
		"DisplayName": currentUser.Fname,
		"Venues": func() (result []interface{}) {
			for _, v := range venues {
				result = append(result, map[string]interface{}{
					"Name": v.Name,
				})
			}
			return
		}(),
	}

	renderHTML(w, "bookVenue.html", viewArgs)
	return
}

func (h *Handler) HandleNewVenueBookingPOST(w http.ResponseWriter, r *http.Request) {
	currentUser := h.currentUser(w, r)
	if currentUser == nil {
		return
	}

	venueKey := "venue"
	venueDateKey := "venueDate"
	venueStKey := "venueSt"
	venueEtKey := "venueEt"
	customerNameKey := "custName"
	customerPhoneKey := "custPhone"

	data, err := postquery.FormValues(r, []string{venueKey, venueDateKey, venueStKey, venueEtKey, customerNameKey, customerPhoneKey})
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
	st, err := time.Parse(time.RFC3339, fmt.Sprintf("%vT%v:00+00:00", data[venueDateKey], data[venueStKey]))
	if err != nil {
		fmt.Printf("[ERROR] unable to parse time, err : %v", err)
		renderJSON(w, http.StatusInternalServerError,
			map[string]string{"error": err500},
		)
		return
	}

	et, err := time.Parse(time.RFC3339, fmt.Sprintf("%vT%v:00+00:00", data[venueDateKey], data[venueEtKey]))
	if err != nil {
		fmt.Printf("[ERROR] unable to parse time, err : %v", err)
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

	if err = h.wDB.BookVenue(data[venueKey], st.String(), et.String(), currentUser.ID, data[customerNameKey], data[customerPhoneKey]); err != nil {
		fmt.Printf("[ERROR] unable to book a venue, err : %v", err)
		renderJSON(w, http.StatusInternalServerError,
			map[string]string{"error": err500},
		)
		return
	}

	renderJSON(w, http.StatusOK, map[string]string{"next": fmt.Sprintf("/user/bookings")})

}

func (h *Handler) HandleAvailableVenuesGET(w http.ResponseWriter, r *http.Request) {
	currentUser := h.currentUser(w, r)
	if currentUser == nil {
		return
	}

	venues, err := h.rDB.GetAvailableVenues()
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
					"Id":    v.ID,
					"Name":  v.Name,
					"Image": v.Image,
				})
			}
			return
		}(),
	}
	renderHTML(w, "availableVenues.html", viewArgs)
	return
}

func (h *Handler) HandleBookedVenuesGET(w http.ResponseWriter, r *http.Request) {
	currentUser := h.currentUser(w, r)
	if currentUser == nil {
		return
	}

	venues, err := h.rDB.GetBookedVenues()
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
					"Id":    v.ID,
					"Name":  v.Name,
					"Image": v.Image,
				})
			}
			return
		}(),
	}
	renderHTML(w, "bookedVenues.html", viewArgs)
	return
}
