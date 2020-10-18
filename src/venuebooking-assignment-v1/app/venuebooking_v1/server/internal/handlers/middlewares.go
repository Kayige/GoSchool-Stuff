package handlers

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/dchest/authcookie"
	"github.com/venuebooking/lib/request"
)

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

func MiddlewareRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		r = r.WithContext(request.AssignRequestID(ctx))
		next.ServeHTTP(w, r)
	})
}

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

func MiddlewareAuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add some common security headers
		w.Header().Add("X-Frame-Options", "SAMEORIGIN")
		w.Header().Add("X-XSS-Protection", "1; mode=block")
		w.Header().Add("X-Content-Type-Options", "nosniff")

		// to check authenticated user
		if cook, err := r.Cookie(cookieForAdmin); err == nil {
			login := authcookie.Login(cook.Value, []byte(cookieSecretAdmin))
			if login == "" {
				http.Redirect(w, r, fmt.Sprintf("/login?continue=%s", r.RequestURI), http.StatusTemporaryRedirect)
				return
			}
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "admin_id", login)))
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/login?continue=%s", r.RequestURI), http.StatusTemporaryRedirect)
		return
	})
}
