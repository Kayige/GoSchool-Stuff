package handlers

import (
	"log"
	"net/http"

	"github.com/venuebooking/lib/request"
)

const (
	err400 = "Invalid Request!"
	err500 = "Internal Server Error!"
)

// Err struct
type Err struct {
	Error string `json:"error"`
}

// recoverError prints the request stack and returns 500(Internal Server Error) status code
func recoverError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("error=internal requestID=%v errorMsg=%v", request.GetRequestID(r.Context()), err)
	w.WriteHeader(http.StatusInternalServerError)
	return
}
