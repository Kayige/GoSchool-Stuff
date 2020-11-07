package db

// Venue struct
type Venue struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

// GetVenueList runs db query and populates into Venue struct
func (r *ReaderDB) GetVenueList() ([]Venue, error) {
	rows, err := r.db.Query("select * from venues;")
	if err != nil {
		return nil, err
	}

	var venues []Venue
	for rows.Next() {
		v := Venue{}
		if err = rows.Scan(&v.ID, &v.Name, &v.Image); err != nil {
			return nil, err
		}
		venues = append(venues, v)
	}

	return venues, nil
}

// GetSearchedVenue takes in name variable and sends query to sql db
func (r *ReaderDB) GetSearchedVenue(name string) ([]Venue, error) {
	query := "select * from venues where name LIKE '%" + name + "%'"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var venues []Venue
	for rows.Next() {
		v := Venue{}
		if err = rows.Scan(&v.ID, &v.Name, &v.Image); err != nil {
			return nil, err
		}
		venues = append(venues, v)
	}

	return venues, nil
}

// GetAvailableVenues checks Venue struct and checks if venue without bookings
func (r *ReaderDB) GetAvailableVenues() ([]Venue, error) {
	rows, err := r.db.Query("select venues.id,venues.name,venues.img from venues left join bookings on venues.name=bookings.v_name where bookings.v_name is NULL;")
	if err != nil {
		return nil, err
	}

	var venues []Venue
	for rows.Next() {
		v := Venue{}
		if err = rows.Scan(&v.ID, &v.Name, &v.Image); err != nil {
			return nil, err
		}
		venues = append(venues, v)
	}

	return venues, nil
}

// GetBookedVenues checks Venue struct and returns venues that are booked.
func (r *ReaderDB) GetBookedVenues() ([]Venue, error) {
	rows, err := r.db.Query("select distinct venues.id,venues.name,venues.img from venues inner join bookings on venues.name=bookings.v_name;")
	if err != nil {
		return nil, err
	}

	var venues []Venue
	for rows.Next() {
		v := Venue{}
		if err = rows.Scan(&v.ID, &v.Name, &v.Image); err != nil {
			return nil, err
		}
		venues = append(venues, v)
	}

	return venues, nil
}

// SaveVenue takes in 2 variables name, img and sends it to MySQL db
func (w *WriterDB) SaveVenue(name, img string) error {
	stmt, err := w.db.Prepare("insert into venues(name,img) values(?,?);")
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(name, img); err != nil {
		return err
	}
	return nil
}
