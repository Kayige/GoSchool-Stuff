package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (api *API) getAllKeys(w http.ResponseWriter, r *http.Request) {
	key, err := api.keyController.GetAllKeys()

	if err != nil {
		response := createResponse(Error, "Internal Server Error", nil)
		sendResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := createResponse(Message, "OK", key)
	sendResponse(w, http.StatusOK, response)
}

func (api *API) createAPIKey(w http.ResponseWriter, r *http.Request) {
	key, err := api.keyController.CreateKey()

	if err != nil {
		response := createResponse(Error, "Internal Server Error", nil)
		sendResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := createResponse(Message, "OK", key)
	sendResponse(w, http.StatusOK, response)
}

func (api *API) deleteAPIKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	err := api.keyController.DeleteKey(key)

	if err != nil {
		response := createResponse(Error, "error deleting key", nil)
		sendResponse(w, http.StatusNotFound, response)
		return
	}

	response := createResponse(Message, "OK", "Deleted")
	sendResponse(w, http.StatusOK, response)
}

func (api *API) getAPIKeyByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	keyResponse, err := api.keyController.GetKeyByID(key)

	if err != nil {
		response := createResponse(Error, "error getting key", nil)
		sendResponse(w, http.StatusNotFound, response)
		return
	}

	response := createResponse(Message, "OK", keyResponse)
	sendResponse(w, http.StatusOK, response)
}

func (api *API) toggleAccess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	keyResponse, err := api.keyController.ToggleAccess(key)

	if err != nil {
		response := createResponse(Error, "error changing status", nil)
		sendResponse(w, http.StatusNotFound, response)
		return
	}

	response := createResponse(Message, "OK", keyResponse)
	sendResponse(w, http.StatusOK, response)
}
