package db

// Booking struct
type Booking struct {
	ID        int64  `json:"id"`
	VenueName string `json:"venueName"`
	St        string `json:"st"`
	Et        string `json:"et"`
	CustName  string `json:"cust_name"`
	CustPhone string `json:"cust_phone"`
}

// BookVenue takes in 6 variables, name, start time, end time, bookedBy(UserID), customer Name, Contact
// Allows UserRoleClient to book a venue from the
func (w *WriterDB) BookVenue(name, st, et string, bookedBy int64, custName, custPhone string) error {
	stmt, err := w.db.Prepare("insert into bookings(v_name,st,et,booked_by,customer,phone) values(?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(name, st, et, bookedBy, custName, custPhone)
	if err != nil {
		return err
	}
	return nil
}

// GetBookingsByVenue checks name input from search venue
func (r *ReaderDB) GetBookingsByVenue(name string) ([]Booking, error) {
	// db query booking id, start time, end time, from table bookings where v_name = name
	rows, err := r.db.Query("select id, st,et from bookings where v_name=?", name)
	if err != nil {
		return nil, err
	}
	// initialize variable
	var bookings []Booking

	// scan rows inside db
	for rows.Next() {
		booking := Booking{VenueName: name}

		if err = rows.Scan(&booking.ID, &booking.St, &booking.Et); err != nil {
			return nil, err
		}

		bookings = append(bookings, booking)
	}
	return bookings, nil
}

// GetBookingsByUser queries the db and returns the Venues booked by referenced User
func (r *ReaderDB) GetBookingsByUser(bookedBy int64) ([]Booking, error) {
	rows, err := r.db.Query("select id,v_name,st,et,customer,phone from bookings where booked_by=?", bookedBy)
	if err != nil {
		return nil, err
	}

	var bookings []Booking

	for rows.Next() {
		booking := Booking{}

		if err = rows.Scan(&booking.ID, &booking.VenueName, &booking.St, &booking.Et, &booking.CustName, &booking.CustPhone); err != nil {
			return nil, err
		}

		bookings = append(bookings, booking)
	}
	return bookings, nil
}

// GetBookingByID queries the database based on the Venue Name and returns the Venue
func (r *ReaderDB) GetBookingByID(bookingID string) (Booking, error) {
	var booking Booking
	rows, err := r.db.Query("select id,v_name,st,et,customer,phone from bookings where id=?", bookingID)
	if err != nil {
		return booking, err
	}

	for rows.Next() {
		if err = rows.Scan(&booking.ID, &booking.VenueName, &booking.St, &booking.Et, &booking.CustName, &booking.CustPhone); err != nil {
			return booking, err
		}
	}
	return booking, nil
}

// UpdateBooking allows the editing of Venue Booking details.
func (w *WriterDB) UpdateBooking(bookingID, name, st, et, custName, custPhone string) error {
	stmt, err := w.db.Prepare("update bookings set v_name=? ,st=? ,et=? ,customer=? ,phone=? where id=?;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(name, st, et, custName, custPhone, bookingID)
	if err != nil {
		return err
	}
	return nil
}
