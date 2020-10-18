package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

const (
	cookieName   = "venue.booking.user"
	cookieSecret = "$P$Bd2WdVjaRR/De58OX2qVu3XA6aiPaf."

	cookieForAdmin    = "venue.booking.admin"
	cookieSecretAdmin = "$P$Bd2WdVjaRR/De89OX2qVu3XA6aiPaf."
)

var Templates *template.Template


// renderHTML function that gets the http response, templateName and viewArgs
func renderHTML(w http.ResponseWriter, templateName string, viewArgs map[string]interface{}) {
	if viewArgs == nil {
		viewArgs = map[string]interface{}{}
	}
	s1 := Templates.Lookup(templateName)
	if s1 == nil {
		http.Error(w, "template not found", http.StatusInternalServerError)
		return
	}
	err := s1.ExecuteTemplate(w, templateName, viewArgs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func renderJSON(w http.ResponseWriter, status int, data interface{}) {
	x, _ := json.MarshalIndent(data, "", "  ")
	log.Printf("status=%v msg=%q", status, string(x))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(x)
}
