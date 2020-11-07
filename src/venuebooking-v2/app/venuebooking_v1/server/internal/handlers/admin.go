package handlers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dchest/authcookie"
	"github.com/venuebooking/app/db"
	"github.com/venuebooking/app/venuebooking_v1/server/internal/config"
	"github.com/venuebooking/lib/crypto"
	"github.com/venuebooking/lib/postquery"
)

// HandleAdminLoginGET returns admin login html page
func (h *Handler) HandleAdminLoginGET(w http.ResponseWriter, r *http.Request) {
	viewArgs := map[string]interface{}{}
	renderHTML(w, "admin-login.html", viewArgs)
	return
}

// HandleAdminLogout destroys admin user cookie
// redirects to admin login page
func (h *Handler) HandleAdminLogout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Domain:  config.Domain(),
		Name:    cookieForAdmin,
		Expires: time.Now(),
		Path:    "/admin",
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/admin/login", http.StatusTemporaryRedirect)
	return
}

// HandleAdminLoginPOST handler verifies admin user credentials
// check for required form parameters
// issue new cookie to admin user if email and password are correct
// redirects admin user to home page
func (h *Handler) HandleAdminLoginPOST(w http.ResponseWriter, r *http.Request) {
	email, err := postquery.RequiredFormParamString(r, "email")
	if err != nil {
		h.badRequestError(w, r, err)
		return
	}
	password, err := postquery.RequiredFormParamString(r, "password")
	if err != nil {
		h.badRequestError(w, r, err)
		return
	}

	user, err := h.rDB.UserByEmail(r.Context(), email, 1)
	if err != nil {
		if _, ok := err.(*db.UserNotFoundError); ok {
			h.badRequestError(w, r, &db.UserNotFoundError{})
			return
		}
		h.internalServerError(w, r, err)
		return

	}

	if user.Role != 1 {
		h.badRequestError(w, r, errors.New("invalid account"))
		return
	}

	// verify hashed password
	if !crypto.PortableHashCheck(password, user.HashedPassword) {
		h.badRequestError(w, r, errors.New("invalid username/password"))
		return
	}

	cookie := &http.Cookie{
		Domain:  config.Domain(),
		Name:    cookieForAdmin,
		Value:   authcookie.NewSinceNow(fmt.Sprintf("%v", user.ID), 24*time.Hour, []byte(cookieSecretAdmin)),
		Expires: time.Now().Add(24 * time.Hour),
		Path:    "/admin",
	}

	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/admin/bookings", http.StatusSeeOther)
	return
}

// HandleAllBookingsGET returns admin booking page
// get all bookings from Database
// pass bookings data to html page
func (h *Handler) HandleAllBookingsGET(w http.ResponseWriter, r *http.Request) {
	// check if admin is already logged in
	// if not logged in then returns to login page
	currentAdmin := h.currentAdmin(w, r)
	if currentAdmin == nil {
		return
	}

	bookings, err := h.rDB.GetBookingsAll()
	if err != nil {
		fmt.Printf("[ERROR] unable to get bookings, err : %v", err)
		renderJSON(w, http.StatusInternalServerError,
			map[string]string{"error": err500},
		)
		return
	}

	viewArgs := map[string]interface{}{
		"Header": map[string]interface{}{
			"DisplayName": currentAdmin.Fname,
		},
		"DisplayName": currentAdmin.Fname,
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
	renderHTML(w, "adminBookings.html", viewArgs)
	return
}

// HandleEditBookingGET returns Edit Booking page
// get venues and booking from database
// pass venues and booking data to html page
func (h *Handler) HandleEditBookingGET(w http.ResponseWriter, r *http.Request) {
	currentAdmin := h.currentAdmin(w, r)
	if currentAdmin == nil {
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
			"DisplayName": currentAdmin.Fname,
		},
		"DisplayName": currentAdmin.Fname,
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
	renderHTML(w, "adminUpdateBooking.html", viewArgs)
	return
}

// HandleBookingPOST update booking
// verify all required form parameters
// update booking details in DB
// redirect to current booking page
func (h *Handler) HandleBookingPOST(w http.ResponseWriter, r *http.Request) {
	currentAdmin := h.currentAdmin(w, r)
	if currentAdmin == nil {
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

	// parse time into the format required by Database
	st, err := time.Parse(time.RFC3339, fmt.Sprintf("%vT%v+00:00", data[venueDateKey], data[venueStKey]))
	if err != nil {
		fmt.Printf("[ERROR] unable to parse time, err : %v", err)
		renderJSON(w, http.StatusInternalServerError,
			map[string]string{"error": err500},
		)
		return
	}

	// parse time into the format required by Database
	et, err := time.Parse(time.RFC3339, fmt.Sprintf("%vT%v+00:00", data[venueDateKey], data[venueEtKey]))
	if err != nil {
		fmt.Printf("[ERROR] unable to parse time, err : %v", err)
		renderJSON(w, http.StatusInternalServerError,
			map[string]string{"error": err500},
		)
		return
	}

	// End Time should be greater than Start Time when booking a venue
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

	renderJSON(w, http.StatusOK, map[string]string{"next": fmt.Sprintf("/admin/bookings/edit?b=%v", data[bookingIDKey])})
}

// HandleAllUsersGET returns admin users html page
// get all users from Database
// pass users data to html page
func (h *Handler) HandleAllUsersGET(w http.ResponseWriter, r *http.Request) {
	currentAdmin := h.currentAdmin(w, r)
	if currentAdmin == nil {
		return
	}

	users, err := h.rDB.GetUsersAll()
	if err != nil {
		fmt.Printf("[ERROR] unable to get list of users, err : %v", err)
		renderJSON(w, http.StatusInternalServerError,
			map[string]string{"error": err500},
		)
		return
	}

	viewArgs := map[string]interface{}{
		"Header": map[string]interface{}{
			"DisplayName": currentAdmin.Fname,
		},
		"DisplayName": currentAdmin.Fname,
		"Users": func() (result []interface{}) {
			for _, user := range users {

				isLoggedIn := "false"
				if user.Session != "" {
					isLoggedIn = "true"
				}
				result = append(result, map[string]interface{}{
					"ID":         user.ID,
					"Fname":      user.Fname,
					"Lname":      user.Lname,
					"Email":      user.Email,
					"CreatedAt":  user.CreatedAt,
					"IsLoggedIn": isLoggedIn,
				})
			}
			return
		}(),
	}
	renderHTML(w, "adminUsers.html", viewArgs)
	return
}

// HandleDelUserGET removes the user from database
// redirects to admin users page
func (h *Handler) HandleDelUserGET(w http.ResponseWriter, r *http.Request) {
	currentAdmin := h.currentAdmin(w, r)
	if currentAdmin == nil {
		return
	}

	err := h.wDB.RemoveUser(r.URL.Query().Get("u"))
	if err != nil {
		fmt.Printf("[ERROR] unable to delete user, err : %v", err)
		renderJSON(w, http.StatusInternalServerError,
			map[string]string{"error": err500},
		)
		return
	}

	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
	return
}

// HandleDelUserSessionGET delete the user session from database
// also delete cookie from browser
func (h *Handler) HandleDelUserSessionGET(w http.ResponseWriter, r *http.Request) {
	currentAdmin := h.currentAdmin(w, r)
	if currentAdmin == nil {
		return
	}

	err := h.wDB.DeleteSession(r.URL.Query().Get("u"))
	if err != nil {
		fmt.Printf("[ERROR] unable to delete user session, err : %v", err)
		renderJSON(w, http.StatusInternalServerError,
			map[string]string{"error": err500},
		)
		return
	}
	cookie := &http.Cookie{
		Domain:  config.Domain(),
		Name:    cookieName,
		Expires: time.Now(),
		Path:    "/user",
	}

	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
	return
}

// HandleVenuesGET returns admin venue list html page
// get all venues from database
// pass venues data to html page
func (h *Handler) HandleVenuesGET(w http.ResponseWriter, r *http.Request) {
	currentAdmin := h.currentAdmin(w, r)
	if currentAdmin == nil {
		return
	}

	venues, err := h.rDB.GetVenueList()
	if err != nil {
		h.internalServerError(w, r, errors.New(err500))
		return
	}

	viewArgs := map[string]interface{}{
		"Header": map[string]interface{}{
			"DisplayName": currentAdmin.Fname,
		},
		"DisplayName": currentAdmin.Fname,
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
	renderHTML(w, "adminVenueList.html", viewArgs)
	return
}

// HandleAdminVenueDetailsGET returns venue details html page
// check if all required query parameters are present
// get Bookings for venue and pass venue data to html page
func (h *Handler) HandleAdminVenueDetailsGET(w http.ResponseWriter, r *http.Request) {
	currentAdmin := h.currentAdmin(w, r)
	if currentAdmin == nil {
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

	// get bookings by venue
	vList, err := h.rDB.GetBookingsByVenue(venueName)
	if err != nil {
		fmt.Printf("[ERROR] unable to get booked venues, err : %v", err)
		renderJSON(w, http.StatusInternalServerError,
			map[string]string{"error": err500},
		)
		return
	}

	viewArgs := map[string]interface{}{
		"Header": map[string]interface{}{
			"DisplayName": currentAdmin.Fname,
		},
		"DisplayName": currentAdmin.Fname,
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

// HandleNewVenueGET returns new venue html page
func (h *Handler) HandleNewVenueGET(w http.ResponseWriter, r *http.Request) {
	currentAdmin := h.currentAdmin(w, r)
	if currentAdmin == nil {
		return
	}

	viewArgs := map[string]interface{}{
		"Header": map[string]interface{}{
			"DisplayName": currentAdmin.Fname,
		},
		"DisplayName": currentAdmin.Fname,
	}
	renderHTML(w, "adminNewVenue.html", viewArgs)
	return
}

// HandleNewVenuePOST creates new venue
// check if all form parameters are present
// save venue image to assets/images
// save venue details to DB
// redirects to admin venues page
func (h *Handler) HandleNewVenuePOST(w http.ResponseWriter, r *http.Request) {
	currentAdmin := h.currentAdmin(w, r)
	if currentAdmin == nil {
		return
	}

	file, fileHeader, err := postquery.GetFile(r, "venueImg")
	if err != nil {
		if !strings.EqualFold(strings.ToLower(err.Error()), "http: no such file") {
			renderJSON(w, http.StatusBadRequest,
				map[string]string{"error": err.Error()},
			)
			return
		}
	}

	venueNameKey := "venueName"

	data, err := postquery.FormValues(r, []string{venueNameKey})
	if err != nil {
		renderJSON(w, http.StatusBadRequest,
			map[string]string{"error": err.Error()},
		)
		return
	}

	if fileHeader != nil {

		ext := strings.ToLower(filepath.Ext(fileHeader.Filename))

		path := fmt.Sprintf("app/venuebooking_v1/server/internal/assets/images/%s%s", data[venueNameKey], ext)

		buf, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Printf("[ERROR] unable to read file, error : %v", err)
			renderJSON(w, http.StatusInternalServerError,
				map[string]string{"error": err500},
			)
			return
		}
		err = ioutil.WriteFile(path, buf, os.ModePerm)
		if err != nil {
			fmt.Printf("[ERROR] unable to write file, error : %v", err)
			renderJSON(w, http.StatusInternalServerError,
				map[string]string{"error": err500},
			)
			return
		}

		if err = h.wDB.SaveVenue(data[venueNameKey], fmt.Sprintf("%v%v", data[venueNameKey], ext)); err != nil {
			fmt.Printf("[ERROR] unable to save venue, error : %v", err)
			renderJSON(w, http.StatusInternalServerError,
				map[string]string{"error": err500},
			)
			return
		}

		http.Redirect(w, r, "/admin/venues", http.StatusSeeOther)
		return
	}

}

// HandleAdminVenueDelGET remove venue from database
// redirects to admin venues page
func (h *Handler) HandleAdminVenueDelGET(w http.ResponseWriter, r *http.Request) {
	currentAdmin := h.currentAdmin(w, r)
	if currentAdmin == nil {
		return
	}

	err := h.wDB.RemoveVenue(r.URL.Query().Get("v"))
	if err != nil {
		fmt.Printf("[ERROR] unable to delete venue, err : %v", err)
		renderJSON(w, http.StatusInternalServerError,
			map[string]string{"error": err500},
		)
		return
	}

	http.Redirect(w, r, "/admin/venues", http.StatusSeeOther)
	return
}
