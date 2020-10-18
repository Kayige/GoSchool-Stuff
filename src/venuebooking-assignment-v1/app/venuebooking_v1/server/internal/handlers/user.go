package handlers

import (
	"fmt"
	"net/http"

	"github.com/venuebooking/lib/postquery"
)

func (h *Handler) HandleProfileGET(w http.ResponseWriter, r *http.Request) {
	currentUser := h.currentUser(w, r)
	if currentUser == nil {
		return
	}

	// viewArgs sends user data and populates to html
	viewArgs := map[string]interface{}{
		"Header": map[string]interface{}{
			"DisplayName": currentUser.Fname,
		},
		"DisplayName": currentUser.Fname,
		"Fn":          currentUser.Fname,
		"Ln":          currentUser.Lname,
		"Em":          currentUser.Email,
	}
	renderHTML(w, "profile.html", viewArgs)
	return
}

func (h *Handler) HandleProfilePOST(w http.ResponseWriter, r *http.Request) {
	currentUser := h.currentUser(w, r)
	if currentUser == nil {
		return
	}

	fnKey := "fname"
	lnKey := "lname"
	emKey := "email"

	data, err := postquery.FormValues(r, []string{fnKey, lnKey, emKey})
	if err != nil {
		renderJSON(w, http.StatusBadRequest,
			map[string]string{"error": err.Error()},
		)
		return
	}

	if data[fnKey] == "" || data[lnKey] == "" || data[emKey] == "" {
		renderJSON(w, http.StatusBadRequest,
			map[string]string{"error": "All Fields are required!"},
		)
		return
	}

	if err := h.wDB.UpdateProfile(currentUser.ID, data[fnKey], data[lnKey], data[emKey]); err != nil {
		fmt.Printf("[ERROR] unable to update user profile info, err : %v", err)
		renderJSON(w, http.StatusInternalServerError,
			map[string]string{"error": err500},
		)
		return
	}
	renderJSON(w, http.StatusOK, map[string]string{"next": fmt.Sprintf("/user/profile")})

}
