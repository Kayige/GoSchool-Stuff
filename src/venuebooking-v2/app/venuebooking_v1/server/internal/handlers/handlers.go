package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/venuebooking/app/db"
)

// Handler struct have dbStr,rDB and wDB fields
// all fields are un-exported cannot be accessible outside the package
type Handler struct {
	dbStr string

	rDB *db.ReaderDB
	wDB *db.WriterDB
}

// NewHandler constructs new Handler with given db connection string
func NewHandler(dbStr string) (Handler, error) {
	rDB, err := db.Reader(dbStr)
	if err != nil {
		return Handler{}, err
	}
	wDB, err := db.Writer(dbStr)
	if err != nil {
		return Handler{}, err
	}
	return Handler{
		dbStr: dbStr,
		rDB:   rDB,
		wDB:   wDB,
	}, nil
}

// currentUser returns the current logged in user
// it checks the field name user_id from the request Context
// if not found returns Unauthorized
// See MiddlewareAuth() in middlewares.go
func (h Handler) currentUser(w http.ResponseWriter, r *http.Request) *db.User {
	userID := r.Context().Value("user_id")
	if userID == nil {
		h.unauthorizedError(w, r)
		return nil
	}
	id, ok := userID.(string)
	if !ok {
		h.internalServerError(w, r, errors.New("user_id not a string"))
		return nil
	}

	user, err := h.rDB.UserByID(r.Context(), id)
	if err != nil {
		// check if error is of type *db.UserNotFoundError
		if _, ok := err.(*db.UserNotFoundError); ok {
			h.badRequestError(w, r, errors.New("user not found"))
			return nil
		}
		h.internalServerError(w, r, err)
		return nil
	}

	// check user role
	if user.Role != db.UserRoleClient {
		h.unauthorizedError(w, r)
		return nil
	}

	log.Printf("currentUserLogin=%q currentUserEmail=%q\n", user.ID, user.Email)
	return &user
}

// currentUser returns the current logged in admin user
// it checks the field name admin_id from the request Context
// if not found returns Unauthorized
// See MiddlewareAuthAdmin() in middlewares.go
func (h Handler) currentAdmin(w http.ResponseWriter, r *http.Request) *db.User {
	userID := r.Context().Value("admin_id")
	if userID == nil {
		h.unauthorizedError(w, r)
		return nil
	}
	id, ok := userID.(string)
	if !ok {
		h.internalServerError(w, r, errors.New("id not a string"))
		return nil
	}

	// get user by id
	user, err := h.rDB.UserByID(r.Context(), id)
	if err != nil {
		// check if error is of type *db.UserNotFoundError
		if _, ok := err.(*db.UserNotFoundError); ok {
			h.badRequestError(w, r, errors.New("Admin user does not exists!"))
			return nil
		}
		h.internalServerError(w, r, err)
		return nil
	}
	// check admin role
	if user.Role != 1 {
		h.unauthorizedError(w, r)
		return nil
	}

	log.Printf("currentUserLogin=%q currentUserEmail=%q\n", user.ID, user.Email)
	return &user
}

// redirect redirects the request to next url
func (h *Handler) redirect(w http.ResponseWriter, r *http.Request, next string) {
	renderJSON(w, http.StatusOK, map[string]string{
		"next": next,
	})
	return
}

// internalServerError writes 500(Internal Server Error) status code to request
func (h *Handler) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)

	return
}

// internalServerError writes 404(Not Found Error) status code to request
func (h *Handler) notFoundError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	return
}

// internalServerError writes 401(Unauthorized Request Error) status code to request
func (h *Handler) unauthorizedError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	return
}

// internalServerError writes 400(Bad Request Error) status code to request
func (h *Handler) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusBadRequest)
	res, _ := json.Marshal(Err{Error: err.Error()})
	w.Write(res)
	return
}
