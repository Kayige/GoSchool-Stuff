package handlers

import (
	"fmt"
	"net/http"
	"net/mail"
	"time"

	"github.com/dchest/authcookie"
	"github.com/venuebooking/app/venuebooking_v1/server/internal/config"
	"github.com/venuebooking/lib/postquery"
)

// HandleUserSignupGET returns signup html form page
func (h *Handler) HandleUserSignupGET(w http.ResponseWriter, r *http.Request) {
	viewArgs := map[string]interface{}{}
	renderHTML(w, "signup.html", viewArgs)
	return
}

// HandleUserSignupPOST register new user
// check if the required form parameters are present and correct
// email address check as well as should not be used by someone else
// redirects to login page
func (h *Handler) HandleUserSignupPOST(w http.ResponseWriter, r *http.Request) {
	fnameKey := "fname"
	lnameKey := "lname"
	emailKey := "email"
	passwordKey := "password"

	// check if required form parameters are present
	data, err := postquery.FormValues(r, []string{emailKey,
		fnameKey, lnameKey, emailKey, passwordKey})
	if err != nil {
		renderJSON(w, http.StatusBadRequest,
			map[string]string{"error": err.Error()},
		)
		return
	}

	if data[emailKey] == "" {
		renderJSON(w, http.StatusBadRequest,
			map[string]string{"error": "Email is required!"},
		)
		return
	}

	emailParsed, err := mail.ParseAddress(data[emailKey])
	if err != nil {
		renderJSON(w, http.StatusBadRequest,
			map[string]string{"error": err.Error()},
		)
		return
	}

	exists, err := h.rDB.IsAlreadyExistsEmail(r.Context(), emailParsed.Address)
	if !exists {
		renderJSON(w, http.StatusBadRequest,
			map[string]string{"error": "Email already exists!"},
		)
		return
	}

	err = h.wDB.CreateUserAccount(r.Context(), data[fnameKey], data[lnameKey], emailParsed.Address, data[passwordKey])
	if err != nil {
		renderJSON(w, http.StatusInternalServerError,
			map[string]string{"error": err500},
		)
		return
	}

	cookie := &http.Cookie{
		Domain: config.Domain(),
		Name:   cookieName,
		Value:  authcookie.NewSinceNow(emailParsed.Address, 24*time.Hour, []byte(cookieSecret)),
		Path:   "/user",
	}

	http.SetCookie(w, cookie)

	renderJSON(w, http.StatusOK, map[string]string{
		"next": fmt.Sprintf("/login"),
	})

	return
}
