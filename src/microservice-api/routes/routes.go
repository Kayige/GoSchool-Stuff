package routes

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"restapi/controller"

	"github.com/gorilla/mux"
)

// API Contains all controller reference and http.Handle
type API struct {
	router           http.Handler
	courseController *controller.Course
	keyController    *controller.Key
	userController   *controller.User
}

// Server interfce implements
type Server interface {
	Router() http.Handler
}

// Router return router from API struct
func (api *API) Router() http.Handler {
	return api.router
}

func initServices() *API {
	c := &controller.Course{}
	k := &controller.Key{}
	u := &controller.User{}

	return &API{
		courseController: c,
		keyController:    k,
		userController:   u,
	}
}

// HomeHandler is for home page of API
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Course API V1.0")
}

func auth(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) != 2 {
			w.Write([]byte("Empty access token. Please set access key in bearer token"))
			return
		} else {
			reqToken = splitToken[1]
			log.Println("Access token: ", reqToken)
			result := controller.AuthenticateKey(reqToken)
			log.Println("it is ", result)
			if result == false {
				w.Write([]byte("Unauthorized access"))
				return
			}
		}

		f(w, r)
	}
}

// New  handdle http endpoinst
func New() Server {
	api := &API{}
	initServices()
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).Methods(http.MethodGet)
	r.HandleFunc("/courses", auth(api.getAllCourses)).Methods(http.MethodGet)
	r.HandleFunc("/course", auth(api.createCourse)).Methods(http.MethodPost)
	r.HandleFunc("/course/{id}", auth(api.getCourseByID)).Methods(http.MethodGet)
	r.HandleFunc("/course/{id}", auth(api.updateCourseByID)).Methods(http.MethodPut)
	r.HandleFunc("/course/{id}", auth(api.deleteCourseByID)).Methods(http.MethodDelete)

	r.HandleFunc("/keys", api.getAllKeys).Methods(http.MethodGet)
	r.HandleFunc("/key", api.createAPIKey).Methods(http.MethodPost)
	r.HandleFunc("/key/{id}", api.deleteAPIKey).Methods(http.MethodDelete)
	r.HandleFunc("/key/{id}", api.getAPIKeyByID).Methods(http.MethodGet)
	r.HandleFunc("/key/{id}", api.toggleAccess).Methods(http.MethodPut)

	r.HandleFunc("/user", auth(api.createUser)).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", auth(api.getUserByID)).Methods(http.MethodGet)
	r.HandleFunc("/users", auth(api.getUsers)).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", auth(api.updateUser)).Methods(http.MethodPut)
	r.HandleFunc("/user/{id}", auth(api.deleteUserByID)).Methods(http.MethodDelete)

	api.router = r
	return api
}
