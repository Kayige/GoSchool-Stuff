package handlers

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/dchest/authcookie"
	"github.com/venuebooking/lib/request"
)

// MiddlewareRecovery recovers from panic if happened in the reqeust stack
// also prints the request stack
func MiddlewareRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				debug.PrintStack()

				recoverError(w, r, fmt.Errorf("%v", err))
				return
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// MiddlewareRequestID assign a unique ID to every Request
func MiddlewareRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		r = r.WithContext(request.AssignRequestID(ctx))
		next.ServeHTTP(w, r)
	})
}

// MiddlewareAuth is an authentication middleware for type user which
// check if user is authenticated or not
func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add some common security headers
		w.Header().Add("X-Frame-Options", "SAMEORIGIN")
		w.Header().Add("X-XSS-Protection", "1; mode=block")
		w.Header().Add("X-Content-Type-Options", "nosniff")

		// to check authenticated user
		if cook, err := r.Cookie(cookieName); err == nil {
			login := authcookie.Login(cook.Value, []byte(cookieSecret))
			if login == "" {
				http.Redirect(w, r, fmt.Sprintf("/login?continue=%s", r.RequestURI), http.StatusTemporaryRedirect)
				return
			}

			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user_id", login)))
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/login?continue=%s", r.RequestURI), http.StatusTemporaryRedirect)
		return
	})
}

// MiddlewareAuthAdmin is an authentication middleware for type admin which
// checks if admin user is authenticated or not
// if authenticated then set admin_id value
func MiddlewareAuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add some common security headers
		w.Header().Add("X-Frame-Options", "SAMEORIGIN")
		w.Header().Add("X-XSS-Protection", "1; mode=block")
		w.Header().Add("X-Content-Type-Options", "nosniff")

		// to check authenticated admin user
		if cook, err := r.Cookie(cookieForAdmin); err == nil {
			login := authcookie.Login(cook.Value, []byte(cookieSecretAdmin))
			if login == "" {
				http.Redirect(w, r, fmt.Sprintf("/admin/login?continue=%s", r.RequestURI), http.StatusTemporaryRedirect)
				return
			}
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "admin_id", login)))
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/admin/login?continue=%s", r.RequestURI), http.StatusTemporaryRedirect)
		return
	})
}
