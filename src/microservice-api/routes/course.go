package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	Error   = "error"
	Message = "message"
)

type response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func createResponse(messageType string, message string, data interface{}) response {
	return response{messageType, message, data}
}

func sendResponse(w http.ResponseWriter, statusCode int, resp response) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(&resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (api *API) createCourse(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&api.courseController); err != nil {
		response := createResponse(Error, "Bad Request", nil)
		sendResponse(w, http.StatusBadRequest, response)
		return
	}

	err := api.courseController.CreateCourse(api.courseController)
	if err != nil {
		response := createResponse(Error, "error creating course", nil)
		sendResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := createResponse(Message, "Course created", nil)
	sendResponse(w, http.StatusCreated, response)

}

func (api *API) getCourseByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	course, err := api.courseController.GetCourseByID(id)

	if err != nil {
		response := createResponse(Error, "error creating course", nil)
		sendResponse(w, http.StatusNotFound, response)
		return
	}

	response := createResponse(Message, "OK", course)
	sendResponse(w, http.StatusOK, response)

}

func (api *API) deleteCourseByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := api.courseController.DeleteCourseByID(id)

	if err != nil {
		response := createResponse(Error, "error deleting course", nil)
		sendResponse(w, http.StatusNotFound, response)
		return
	}

	response := createResponse(Message, "OK", "Deleted")
	sendResponse(w, http.StatusOK, response)

}

func (api *API) getAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := api.courseController.GetAllCourses()

	if err != nil {
		response := createResponse(Error, "Internal Server Error", nil)
		sendResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := createResponse(Message, "OK", courses)
	sendResponse(w, http.StatusOK, response)
}

func (api *API) updateCourseByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := json.NewDecoder(r.Body).Decode(&api.courseController); err != nil {
		response := createResponse(Error, "Invalid structure", nil)
		sendResponse(w, http.StatusBadRequest, response)
		return
	}

	course, err := api.courseController.UpdateCourse(id, *api.courseController)

	if err != nil {
		response := createResponse(Error, "Internal Server Error", nil)
		sendResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := createResponse(Message, "OK", course)
	sendResponse(w, http.StatusOK, response)
}
