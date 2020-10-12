package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func (v Venue) newVenue() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Panic("Error while reading name:", err)
	}

	var capacity int
	fmt.Print("Enter capacity: ")
	_, err = fmt.Scanf("%d\n", &capacity)
	if err != nil {
		log.Panic("Error while reading capacity:", err)
	}

	fmt.Print("Enter type: ")
	roomType, err := reader.ReadString('\n')
	if err != nil {
		log.Panic("Error while reading type:", err)
	}

	var availableDates []DateTime

	fmt.Println("You'll now begin entering available dates for the venue, if you want to stop entering new dates type exit when asked for a date")
	for {
		var date string
		fmt.Print("Enter available date (use format YYYY-MM-DD): ")
		_, err = fmt.Scanf("%s\n", &date)
		if err != nil {
			log.Panic("Error while reading date:", err)
		}

		if strings.ToLower(date) == "exit" {
			break
		}
		dateTime, _ := time.Parse("2006-01-02", date)

		var clock string
		fmt.Print("Enter time (use format 3:04PM): ")
		_, err = fmt.Scanf("%s\n", &clock)
		clockTime, _ := time.Parse("3:04PM", clock)
		if err != nil {
			log.Panic("Error while reading clock time:", err)
		}

		newDate := DateTime{
			Date: dateTime,
			Time: clockTime,
		}
		availableDates = append(availableDates, newDate)
	}

	newVenue := Venue{
		Name:           name,
		Capacity:       capacity,
		Type:           roomType,
		AvailableDates: availableDates,
	}

	allVenues = append(allVenues, newVenue)
}

// SearchVenue Function to search for venue
func (v Venue) SearchVenue() {
	var venuesFound []Venue
	var searchData string
	var tmp string

	fmt.Printf("Enter the terms you want to search using the format date,time,capacity,type\nIf you dont need a term use '-'\nFor example:\n2020-10-10,-,300,-\n:")
	fmt.Scanf("%s", &searchData)

	for _, v := range allVenues {
		if v.DataInVenue(searchData, v) {
			venuesFound = append(venuesFound, v)
		}
	}

	if len(venuesFound) == 0 {
		fmt.Print("NO VENUE HAS BEEN FOUND...")
		fmt.Scanf("\n%s", &tmp)
		return
	}

	fmt.Println("VENUES FOUND")
	for venueIndex, v := range venuesFound {
		fmt.Printf("-------------------------------------------------\nVENUE %d:\n", venueIndex)
		fmt.Printf("Name: %s\nCapacity: %d\nType: %s\nAvailable dates:\n", v.Name, v.Capacity, v.Type)
		for _, d := range v.AvailableDates {
			fmt.Printf("%s %s\n", d.Date.Format("2006-01-02"), d.Time.Format("3:04PM"))
		}
		fmt.Printf("\n-------------------------------------------------\n\n")
	}
	fmt.Scanf("\n%s", &tmp)
}

// DataInVenue function
func (v Venue) DataInVenue(searchData string, venue Venue) bool {
	matches := 0
	data := strings.Split(searchData, ",")

	for dataIndex, d := range data {
		if d == "-" {
			matches++
		} else {
			if dataIndex == 0 {
				date, _ := time.Parse("2006-01-02", data[0])
				for _, a := range venue.AvailableDates {
					if date == a.Date {
						matches++
					}
				}
			} else if dataIndex == 1 {
				time, _ := time.Parse("3:04PM", data[1])
				for _, a := range venue.AvailableDates {
					if time == a.Time {
						matches++
					}
				}
			} else if dataIndex == 2 {
				capacity, _ := strconv.Atoi(data[2])
				if capacity == venue.Capacity {
					matches++
				}
			} else if dataIndex == 3 {
				if data[3] == venue.Type {
					matches++
				}
			}
		}
	}
	// roomType :=

	if matches == 4 {
		return true
	}

	return false
}

func (v Venue) browseVenues(ch chan<- int) {
	// Goes through all the venues, if s is entered then the index of the selected venue will be returned, if no venue was selected it will return -1
	if len(allVenues) == 0 {
		fmt.Print("NO VENUE HAS BEEN ADDED...")
		var tmp string
		fmt.Scanf("%s\n", &tmp)
		return
	}

	for venueIndex, v := range allVenues {
		fmt.Printf("-------------------------------------------------\nVENUE %d:\n", venueIndex)
		fmt.Printf("Name: %s\nCapacity: %d\nType: %s\nAvailable dates:\n", v.Name, v.Capacity, v.Type)
		for _, d := range v.AvailableDates {
			fmt.Printf("%s %s\n", d.Date.Format("2006-01-02"), d.Time.Format("3:04PM"))
		}
		fmt.Printf("\n-------------------------------------------------\nPress enter to see next venue...\n: ")
		var next string
		fmt.Scanf("%s\n", &next)
		if ch != nil {
			if next == "s" {
				ch <- venueIndex
				return
			}
		}
	}
}
