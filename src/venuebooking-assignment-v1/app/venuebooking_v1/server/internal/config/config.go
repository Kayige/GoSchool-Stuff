package config

import (
	"os"
)

// HTTPAddr function that returns the port number
func HTTPAddr() string {
	addr := os.Getenv("HTTP_ADDR")
	if addr != "" {
		return addr
	}
	return ":8080"
}

// AssetsDir function that returns the path to the image & js
func AssetsDir() string {
	views := os.Getenv("ASSETS")
	if views != "" {
		return views
	}
	return "app/venuebooking_v1/server/internal/assets/"
}

// ViewsDir function that returns the path to the html files
func ViewsDir() string {
	views := os.Getenv("VIEWS")
	if views != "" {
		return views
	}
	return "app/venuebooking_v1/server/internal/views"
}

// DBConnectionString that returns the MySQL database connection
func DBConnectionString() string {
	db := os.Getenv("DB")
	if db != "" {
		return db
	}

	// please change this line to ensure sql database connection
	return "user:password@tcp(localhost:3306)/venue_booking"
}

// Domain function that sets the domain to localhost
func Domain() string {
	domain := os.Getenv("DOMAIN")
	if domain != "" {
		return domain
	}
	return "localhost"
}
