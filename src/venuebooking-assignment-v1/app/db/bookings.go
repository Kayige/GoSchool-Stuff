package db

// Booking struct
type Booking struct {
	Id        int64  `json:"id"`
	VenueName string `json:"venueName"`
	St        string `json:"st"`
	Et        string `json:"et"`
	CustName  string `json:"cust_name"`
	CustPhone string `json:"cust_phone"`
}

// BookVenue takes in 6 variables, name, start time, end time, bookedBy(UserID), customer Name, Contact
// Prepares SQL query statement and writes to SQL database
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

// GetBookingsByVenue takes a string input
// Sends a SQL query via string input checks each booking
// returns the searched bookings
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

		if err = rows.Scan(&booking.Id, &booking.St, &booking.Et); err != nil {
			return nil, err
		}

		bookings = append(bookings, booking)
	}
	return bookings, nil
}

// GetBookingsByUser checks for uuid
// Sends a query to filter out uuid
// returns the booking with uuid
func (r *ReaderDB) GetBookingsByUser(booked_by int64) ([]Booking, error) {
	rows, err := r.db.Query("select id,v_name,st,et,customer,phone from bookings where booked_by=?", booked_by)
	if err != nil {
		return nil, err
	}

	var bookings []Booking

	for rows.Next() {
		booking := Booking{}

		if err = rows.Scan(&booking.Id, &booking.VenueName, &booking.St, &booking.Et, &booking.CustName, &booking.CustPhone); err != nil {
			return nil, err
		}

		bookings = append(bookings, booking)
	}
	return bookings, nil
}

// GetBookingById takes in booking_id string
// checks the booking array
// returns booking based on booking_id
func (r *ReaderDB) GetBookingById(booking_id string) (Booking, error) {
	var booking Booking
	rows, err := r.db.Query("select id,v_name,st,et,customer,phone from bookings where id=?", booking_id)
	if err != nil {
		return booking, err
	}

	for rows.Next() {
		if err = rows.Scan(&booking.Id, &booking.VenueName, &booking.St, &booking.Et, &booking.CustName, &booking.CustPhone); err != nil {
			return booking, err
		}
	}
	return booking, nil
}

// UpdateBooking takes in 6 parameters
// SQL query updates the booking table
func (w *WriterDB) UpdateBooking(bookingId, name, st, et, custName, custPhone string) error {
	stmt, err := w.db.Prepare("update bookings set v_name=? ,st=? ,et=? ,customer=? ,phone=? where id=?;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(name, st, et, custName, custPhone, bookingId)
	if err != nil {
		return err
	}
	return nil
}
