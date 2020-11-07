package config

import (
	"os"
)

// HTTPAddr returns the http port number
func HTTPAddr() string {
	addr := os.Getenv("HTTP_ADDR")
	if addr != "" {
		return addr
	}
	return ":8080"
}

// AssetsDir returns the path to the image & js
func AssetsDir() string {
	views := os.Getenv("ASSETS")
	if views != "" {
		return views
	}
	return "app/venuebooking_v1/server/internal/assets/"
}

// ViewsDir returns the path to the html files
func ViewsDir() string {
	views := os.Getenv("VIEWS")
	if views != "" {
		return views
	}
	return "app/venuebooking_v1/server/internal/views"
}

// DBConnectionString returns the MySQL database connection
func DBConnectionString() string {
	db := os.Getenv("DB")
	if db != "" {
		return db
	}

	// please change this line to ensure sql database connection
	return "user:password@tcp(localhost:3306)/venue_booking"
}

// Domain returns the domain of application
func Domain() string {
	domain := os.Getenv("DOMAIN")
	if domain != "" {
		return domain
	}
	return "localhost"
}
