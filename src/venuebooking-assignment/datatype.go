package main

import "time"

// VenueInterface to store the methods for reusability
type VenueInterface interface {

	// Methods
	newVenue()
	browseVenues(ch chan<- int)
	listBookings(ch <-chan int) (int, int)
	DataInVenue(searchData string, venue Venue) bool
	SearchVenue()
	BookVenue(ch <-chan int)
	editBooking(venueIndex, bookingIndex int)
	menu(admin bool)
}

// DateTime struct to store Date & Time
type DateTime struct {
	Date time.Time
	Time time.Time
}

// Reservation struct to store the Name of the Person ref to DateTime Struct
type Reservation struct {
	Name string
	Date DateTime
}

// Venue Struct to store data of Venue
type Venue struct {
	Name           string
	Capacity       int
	Type           string
	AvailableDates []DateTime
	Reservations   []Reservation
}
