package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dchest/authcookie"
	"github.com/venuebooking/app/db"
	"github.com/venuebooking/app/venuebooking_v1/server/internal/config"
	"github.com/venuebooking/lib/crypto"
	"github.com/venuebooking/lib/postquery"
)

// HandleUserLoginGET handler returns html login page
func (h *Handler) HandleUserLoginGET(w http.ResponseWriter, r *http.Request) {
	viewArgs := map[string]interface{}{}
	renderHTML(w, "login.html", viewArgs)
	return
}

// HandleUserLoginPOST handler verifies user credentials
// check for required form parameters
// issue new cookie to user if email and password are correct
// redirects user to home page
func (h *Handler) HandleUserLoginPOST(w http.ResponseWriter, r *http.Request) {
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

	user, err := h.rDB.UserByEmail(r.Context(), email, 2)
	if err != nil {
		// check if err is of type *db.UserNotFoundError
		if _, ok := err.(*db.UserNotFoundError); ok {
			h.badRequestError(w, r, &db.UserNotFoundError{})
			return
		}
		h.internalServerError(w, r, err)
		return

	}
	// check user role defined in the db pkg
	if user.Role != db.UserRoleClient {
		h.badRequestError(w, r, errors.New("invalid account"))
		return
	}
	// check password hash
	if !crypto.PortableHashCheck(password, user.HashedPassword) {
		h.badRequestError(w, r, errors.New("invalid username/password"))
		return
	}

	// send session cookie
	sessionToken := authcookie.NewSinceNow(fmt.Sprintf("%v", user.ID), 24*time.Hour, []byte(cookieSecret))

	if err = h.wDB.UpdateSession(user.ID, sessionToken); err != nil {
		h.internalServerError(w, r, errors.New(err500))
		return
	}

	cookie := &http.Cookie{
		Domain:  config.Domain(),
		Name:    cookieName,
		Value:   sessionToken,
		Expires: time.Now().Add(24 * time.Hour),
		Path:    "/user",
	}

	http.SetCookie(w, cookie)

	// redirect to user HOME page on successful authentication
	h.redirect(w, r, "/user/venues")
	return
}

// HandleUserLogout destroys user cookie
// redirects to login page
func (h *Handler) HandleUserLogout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Domain:  config.Domain(),
		Name:    cookieName,
		Expires: time.Now(),
		Path:    "/user",
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	return
}
