package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (api *API) createUser(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&api.userController); err != nil {
		response := createResponse(Error, "Error in struct user", nil)
		sendResponse(w, http.StatusBadRequest, response)
		return
	}

	if err := api.userController.CreateUser(*api.userController); err != nil {
		response := createResponse(Error, "Error creating user", nil)
		sendResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := createResponse(Message, "OK", "User created")
	sendResponse(w, http.StatusOK, response)

}

func (api *API) getUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	user, err := api.userController.GetUserByID(id)

	if err != nil {
		response := createResponse(Error, "User not found", nil)
		sendResponse(w, http.StatusNotFound, response)
		return
	}

	response := createResponse(Message, "OK", &user)
	sendResponse(w, http.StatusOK, response)
}

func (api *API) getUsers(w http.ResponseWriter, r *http.Request) {

	users, err := api.userController.GetUsers()

	if err != nil {
		response := createResponse(Error, "Internal Server Error", nil)
		sendResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := createResponse(Message, "OK", &users)
	sendResponse(w, http.StatusOK, response)
}

func (api *API) updateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := json.NewDecoder(r.Body).Decode(&api.userController); err != nil {
		response := createResponse(Error, "Invalid structure", nil)
		sendResponse(w, http.StatusBadRequest, response)
		return
	}

	user, err := api.userController.UpdateUser(id, *api.userController)

	if err != nil {
		response := createResponse(Error, "Internal Server Error", nil)
		sendResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := createResponse(Message, "OK", &user)
	sendResponse(w, http.StatusOK, response)

}

func (api *API) deleteUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := api.userController.DeleteUserByID(id)

	if err != nil {
		response := createResponse(Error, "error deleting course", nil)
		sendResponse(w, http.StatusNotFound, response)
		return
	}

	response := createResponse(Message, "OK", "Deleted")
	sendResponse(w, http.StatusOK, response)

}
