package server

import (
	"net/http"

	"github.com/venuebooking/app/venuebooking_v1/server/internal/handlers"

	"github.com/venuebooking/app/venuebooking_v1/server/internal/config"

	"github.com/gorilla/mux"
)

// initRoutes() creates new handler
// initialize mux router and binds handler functions to the relevant routes
// returns router
func initRoutes() (http.Handler, error) {
	handler, err := handlers.NewHandler(config.DBConnectionString())
	if err != nil {
		return nil, err
	}

	r := mux.NewRouter()
	// use paths as defined in the r.HandleFunc
	r.StrictSlash(true)

	r.Use(handlers.MiddlewareRequestID, handlers.MiddlewareRecovery)
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(config.AssetsDir()))))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
		return
	}).Methods("GET")

	// user routes
	r.HandleFunc("/login", handler.HandleUserLoginGET).Methods("GET")
	r.HandleFunc("/login", handler.HandleUserLoginPOST).Methods("POST")
	r.HandleFunc("/logout", handler.HandleUserLogout).Methods("GET")

	r.HandleFunc("/signup", handler.HandleUserSignupGET).Methods("GET")
	r.HandleFunc("/signup", handler.HandleUserSignupPOST).Methods("POST")

	userRouter := r.PathPrefix("/user/").Subrouter()

	// use auth middleware for authenticated requests
	userRouter.Use(handlers.MiddlewareAuth)

	userRouter.HandleFunc("/profile", handler.HandleProfileGET).Methods("GET")
	userRouter.HandleFunc("/profile", handler.HandleProfilePOST).Methods("POST")

	userRouter.HandleFunc("/bookings", handler.HandleMyBookingsGET).Methods("GET")
	userRouter.HandleFunc("/bookings/edit", handler.HandleEditMyBookingGET).Methods("GET")
	userRouter.HandleFunc("/bookings", handler.HandleMyBookingsPOST).Methods("POST")

	userRouter.HandleFunc("/venues", handler.HandleVenueListGET).Methods("GET")
	userRouter.HandleFunc("/venue/details", handler.HandleVenueDetailsGET).Methods("GET")

	userRouter.HandleFunc("/venue/new", handler.HandleNewVenueBookingGET).Methods("GET")
	userRouter.HandleFunc("/venue/new", handler.HandleNewVenueBookingPOST).Methods("POST")

	userRouter.HandleFunc("/venue/search", handler.HandleSearchVenueGET).Methods("GET")

	userRouter.HandleFunc("/venue/available", handler.HandleAvailableVenuesGET).Methods("GET")
	userRouter.HandleFunc("/venue/booked", handler.HandleBookedVenuesGET).Methods("GET")

	// admin routes
	r.HandleFunc("/admin/login", handler.HandleAdminLoginGET).Methods("GET")
	r.HandleFunc("/admin/login", handler.HandleAdminLoginPOST).Methods("POST")
	r.HandleFunc("/admin/logout", handler.HandleAdminLogout).Methods("GET")

	adminRouter := r.PathPrefix("/admin/").Subrouter()

	// use auth middleware for authenticated requests
	adminRouter.Use(handlers.MiddlewareAuthAdmin)

	adminRouter.HandleFunc("/bookings", handler.HandleAllBookingsGET).Methods("GET")
	adminRouter.HandleFunc("/bookings/edit", handler.HandleEditBookingGET).Methods("GET")
	adminRouter.HandleFunc("/bookings", handler.HandleBookingPOST).Methods("POST")

	adminRouter.HandleFunc("/users", handler.HandleAllUsersGET).Methods("GET")
	adminRouter.HandleFunc("/users/del", handler.HandleDelUserGET).Methods("GET")
	adminRouter.HandleFunc("/users/session/del", handler.HandleDelUserSessionGET).Methods("GET")

	adminRouter.HandleFunc("/venues", handler.HandleVenuesGET).Methods("GET")
	adminRouter.HandleFunc("/venues/details", handler.HandleAdminVenueDetailsGET).Methods("GET")
	adminRouter.HandleFunc("/venues/del", handler.HandleAdminVenueDelGET).Methods("GET")
	adminRouter.HandleFunc("/venues/new", handler.HandleNewVenueGET).Methods("GET")
	adminRouter.HandleFunc("/venues/new", handler.HandleNewVenuePOST).Methods("POST")

	return r, nil
}
