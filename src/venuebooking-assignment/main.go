package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var clear map[string]func()
var venueCh chan Venue
var allVenues []Venue

// CreateMockData for testing creates 10 Input Data
func CreateMockData() {
	for i := 0; i < 10; i++ {
		var availableDates []DateTime
		for i := 0; i < 3; i++ {
			dateTime, _ := time.Parse("2006-01-02", "2020-10-10")
			clockTime, _ := time.Parse("3:04PM", "11:11PM")
			newDate := DateTime{
				Date: dateTime,
				Time: clockTime,
			}
			availableDates = append(availableDates, newDate)
		}

		name := fmt.Sprintf("Test %d", i)

		newVenue := Venue{
			Name:           name,
			Capacity:       i,
			Type:           "Hall room",
			AvailableDates: availableDates,
		}

		allVenues = append(allVenues, newVenue)
	}
}
func init() {
	CreateMockData()
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// CreateMockData function to
func (v *Venue) CreateMockData() {
	for i := 0; i < 10; i++ {
		var availableDates []DateTime
		for i := 0; i < 3; i++ {
			dateTime, _ := time.Parse("2006-01-02", "2020-10-10")
			clockTime, _ := time.Parse("3:04PM", "11:11PM")
			newDate := DateTime{
				Date: dateTime,
				Time: clockTime,
			}
			availableDates = append(availableDates, newDate)
		}

		name := fmt.Sprintf("Test %d", i)

		newVenue := Venue{
			Name:           name,
			Capacity:       i,
			Type:           "Hall room",
			AvailableDates: availableDates,
		}

		allVenues = append(allVenues, newVenue)
	}
}

func (v *Venue) init() {
	v.CreateMockData()
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// CallClear Function to test Error
func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		log.Panic("Your platform is unsupported! Can't clear terminal :/")
	}
}

func (v Venue) menu(admin bool) {
	var option int
	fmt.Println("Please select an option by selecting a number from the menu:")
	fmt.Println("(1) Browse venue")
	fmt.Println("(2) Search venue")
	fmt.Println("(3) Book venue")
	fmt.Println("(4) Edit venue")
	if admin {
		fmt.Println("(5) Add new venue")
		fmt.Println("(6) List bookings")
	}

	fmt.Printf("Option: ")
	fmt.Scanf("%d\n", &option)

	if option > -1 && option < 7 {
		if option == 1 {
			v.browseVenues(nil)
		} else if option == 2 {
			v.SearchVenue()
		} else if option == 3 {
			ch := make(chan int)
			fmt.Printf("TO SELECT A VENUE TO BOOK ENTER THE LETTER S\n\n")
			go v.browseVenues(ch)
			v.BookVenue(ch)
		} else if option == 4 {
			fmt.Printf("TO SELECT A VENUE ENTER THE LETTER S\n\n")
			ch := make(chan int)
			go v.browseVenues(ch)
			fmt.Printf("TO SELECT A BOOKING TO EDIT ENTER THE LETTER S\n\n")
			venueIndex, bookingIndex := v.listBookings(ch)
			v.editBooking(venueIndex, bookingIndex)
		} else if admin {
			if option == 5 {
				v.newVenue()
			} else if option == 6 {
				ch := make(chan int)
				go v.browseVenues(ch)
				v.listBookings(ch)
			}
		} else {
			fmt.Println("Sorry, you are not admin")
		}
	}
}

func main() {
	var v VenueInterface
	v = Venue{}
	for {
		CallClear()
		v.menu(true)
	}
}
