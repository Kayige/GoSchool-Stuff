package db

func (r *ReaderDB) GetBookingsAll() ([]Booking, error) {
	rows, err := r.db.Query("select id,v_name,st,et,customer,phone from bookings;")
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

func (r *ReaderDB) GetUsersAll() ([]User, error) {
	rows, err := r.db.Query("select id,fname,lname,email,created_at,session from users where role=2;")
	if err != nil {
		return nil, err
	}

	var users []User

	for rows.Next() {
		user := User{}

		if err = rows.Scan(&user.ID, &user.Fname, &user.Lname, &user.Email, &user.CreatedAt, &user.Session); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func (w *WriterDB) RemoveUser(userID string) error {
	stmt, err := w.db.Prepare("delete from users where id=? and role=2;")
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(userID); err != nil {
		return err
	}

	return nil
}

func (w *WriterDB) RemoveVenue(venueID string) error {
	stmt, err := w.db.Prepare("delete from venues where id=?;")
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(venueID); err != nil {
		return err
	}

	return nil
}