package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/venuebooking/app/db"
)

type Handler struct {
	dbStr string

	rDB *db.ReaderDB
	wDB *db.WriterDB
}

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
		if _, ok := err.(*db.UserNotFoundError); ok {
			h.badRequestError(w, r, errors.New("user not found"))
			return nil
		}
		h.internalServerError(w, r, err)
		return nil
	}
	if user.Role != db.UserRoleClient {
		h.unauthorizedError(w, r)
		return nil
	}

	log.Printf("currentUserLogin=%q currentUserEmail=%q\n", user.ID, user.Email)
	return &user
}

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

	user, err := h.rDB.UserByID(r.Context(), id)
	if err != nil {
		if _, ok := err.(*db.UserNotFoundError); ok {
			h.badRequestError(w, r, errors.New("Admin user does not exists!"))
			return nil
		}
		h.internalServerError(w, r, err)
		return nil
	}
	if user.Role != 1 {
		h.unauthorizedError(w, r)
		return nil
	}

	log.Printf("currentUserLogin=%q currentUserEmail=%q\n", user.ID, user.Email)
	return &user
}

func (h *Handler) redirect(w http.ResponseWriter, r *http.Request, next string) {
	renderJSON(w, http.StatusOK, map[string]string{
		"next": next,
	})
	return
}

func (h *Handler) optionalQueryParamUint(w http.ResponseWriter, r *http.Request, name string) *uint64 {
	valueStr := r.URL.Query().Get(name)
	value := uint64(0)
	if valueStr != "" {
		var err error
		value, err = strconv.ParseUint(valueStr, 10, 64)
		if err != nil {
			h.badRequestError(w, r, fmt.Errorf("%v is invalid", name))
			return nil
		}
		return &value
	}
	return &value
}

func (h *Handler) requiredQueryParamUint(w http.ResponseWriter, r *http.Request, name string) *uint64 {
	valueStr := r.URL.Query().Get(name)
	if valueStr == "" {
		h.badRequestError(w, r, fmt.Errorf("%v is required", name))
		return nil
	}
	value, err := strconv.ParseUint(valueStr, 10, 64)
	if err != nil {
		h.badRequestError(w, r, fmt.Errorf("%v is invalid", name))
		return nil
	}
	return &value
}

func (h *Handler) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)

	return
}

func (h *Handler) notFoundError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	return
}

func (h *Handler) unauthorizedError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	return
}

func (h *Handler) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusBadRequest)
	res, _ := json.Marshal(Err{Error: err.Error()})
	w.Write(res)
	return
}
