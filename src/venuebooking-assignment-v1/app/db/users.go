package db

import (
	"context"

	"github.com/venuebooking/lib/crypto"
)

// UserRole variable int
type UserRole int

const (
	// UserRoleAdmin Set to 1
	UserRoleAdmin UserRole = 1
	// UserRoleClient Set to 2
	UserRoleClient = 2
)

type (
	// User struct to parse
	User struct {
		ID             int64  `json:"id"`
		Fname          string `json:"fname"`
		Lname          string `json:"lname"`
		Email          string `json:"email"`
		HashedPassword string `json:"password"`
		CreatedAt      string `json:"created_at"`
		Role           int    `json:"role"`
		Session        string `json:"-"`
	}
)

// UserByEmail receives incoming request and verifies User with database query by email and role
func (r *ReaderDB) UserByEmail(ctx context.Context, email string, role int) (User, error) {
	var user User
	rows, err := r.db.Query("select * from users where email=? and role=?;", email, role)
	if err != nil {
		return user, err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Fname, &user.Lname, &user.Email, &user.HashedPassword, &user.CreatedAt, &user.Role, &user.Session)
		if err != nil {
			return user, err
		}
	}

	if user.Email == "" {
		return user, &UserNotFoundError{}
	}

	return user, nil
}

// UserByID receives incoming request and verifies User with database query by ID
func (r *ReaderDB) UserByID(ctx context.Context, userID string) (User, error) {
	var user User
	rows, err := r.db.Query("select * from users where id=?", userID)
	if err != nil {
		return user, err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Fname, &user.Lname, &user.Email, &user.HashedPassword, &user.CreatedAt, &user.Role, &user.Session)
		if err != nil {
			return user, err
		}
	}

	if user.Email == "" {
		return user, &UserNotFoundError{}
	}

	return user, nil
}

// UpdateProfile receives UserID, firstName, lastName, email and POST to MySQL db to update the entries
func (w *WriterDB) UpdateProfile(id int64, fn, ln, em string) error {
	stmt, err := w.db.Prepare("update users set fname = ? , lname = ? , email = ? where id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(fn, ln, em, id)
	if err != nil {
		return err
	}

	return nil
}

// IsAlreadyExistsEmail checks whether user email is true or false
func (r *ReaderDB) IsAlreadyExistsEmail(ctx context.Context, email string) (bool, error) {
	_, err := r.UserByEmail(ctx, email, 2)
	if _, ok := err.(*UserNotFoundError); ok {
		return true, nil
	}
	return false, err

}

// CreateUserAccount Function that takes in 4 variables, first name, last name, email, password
// Prepares variables to be inserted to SQL DB
func (w *WriterDB) CreateUserAccount(ctx context.Context, fname, lname, email, password string) error {
	stmt, err := w.db.Prepare("insert into users(fname, lname, email, password, session) VALUES(?,?,?,?,?)")
	if err != nil {
		return err
	}
	// uses Crypto library to generate hash to mask password
	_, err = stmt.Exec(fname, lname, email, crypto.CryptPrivate(password, crypto.CRYPT_SETTING), "")
	return err
}

// UpdateSession inserts sessionToken to userID upon login
func (w *WriterDB) UpdateSession(userID int64, sessionToken string) error {
	stmt, err := w.db.Prepare("update users set session=? where id=?")
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(sessionToken, userID); err != nil {
		return err
	}
	return nil
}

// DeleteSession removes user session detail from database
func (r *ReaderDB) DeleteSession(userID string) error {
	stmt, err := r.db.Prepare("update users set session='' where id=?;")
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(userID); err != nil {
		return err
	}

	return nil
}
