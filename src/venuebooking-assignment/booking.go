package main

import "fmt"

func (v Venue) listBookings(ch <-chan int) (int, int) {
	index := <-ch
	if len(allVenues[index].Reservations) == 0 {
		fmt.Print("\nNO RESERVATION HAS BEEN DONE TO THIS VENUE...")
		var tmp string
		fmt.Scanf("%s\n", &tmp)
		return -1, -1
	}
	// Goes through all the bookings of a venue, if s is entered then the index of the selected booking will be returned, if no booking was selected it will return -1
	for bookingIndex, r := range allVenues[index].Reservations {
		fmt.Printf("-------------------------------------------------\nBOOKING %d:\n", bookingIndex)
		fmt.Printf("Reservee: %s\nDate: %s\nTime: %s", r.Name, r.Date.Date.Format("2006-01-02"), r.Date.Time.Format("3:04PM"))
		fmt.Printf("\n-------------------------------------------------\nPress enter to see next booking...\n: ")
		var next string
		fmt.Scanf("%s\n", &next)
		CallClear()
		if next == "s" {
			return index, bookingIndex
		}
	}
	return -1, -1
}

// BookVenue function for Making reservation
func (v Venue) BookVenue(ch <-chan int) {
	index := <-ch
	var name string
	var tmp string
	var next string
	fmt.Printf("Enter the name of the reserve: ")
	fmt.Scanf("%s", &name)
	fmt.Scanf("%s", &tmp)

	for i, d := range allVenues[index].AvailableDates {
		fmt.Printf("\n-------------------------------------------------\nDATE %d:\n", i)
		fmt.Printf("%s %s\n", d.Date.Format("2006-01-02"), d.Time.Format("3:04PM"))
		fmt.Printf("-------------------------------------------------\nPress enter to see next date...\n: ")
		fmt.Scanf("%s", &next)
		if next == "s" {
			newReservation := Reservation{
				Name: name,
				Date: d,
			}
			allVenues[index].Reservations = append(allVenues[index].Reservations, newReservation)
			fmt.Println("VENUE BOOKED SUCCESFULLY...")
			fmt.Scanf("\n%s", &tmp)
			return
		}
	}

	fmt.Scanf("%s", &tmp)
}

func (v Venue) editBooking(venueIndex, bookingIndex int) {
	var next string
	var name string
	var tmp string
	fmt.Printf("Enter a new name for the reserve: ")
	fmt.Scanf("%s", &name)
	fmt.Scanf("%s", &tmp)

	allVenues[venueIndex].Reservations[bookingIndex].Name = name

	for i, d := range allVenues[venueIndex].AvailableDates {
		fmt.Printf("\n-------------------------------------------------\nDATE %d:\n", i)
		fmt.Printf("%s %s\n", d.Date.Format("2006-01-02"), d.Time.Format("3:04PM"))
		fmt.Printf("-------------------------------------------------\nPress enter to see next date...\n: ")
		fmt.Scanf("%s", &next)
		if next == "s" {
			allVenues[venueIndex].Reservations[bookingIndex].Date = d

			fmt.Println("BOOKING EDITED SUCCESSFULLY...")
			fmt.Scanf("\n%s", &tmp)
			return
		}
	}

	fmt.Scanf("%s", &tmp)
}
