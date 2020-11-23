package controller

import "restapi/database"

// User Represent and inherit from base Model
type User struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Age       int    `json:"age,omitempty"`
}

// CreateUser creates user in db, return err if not ok
func (u *User) CreateUser(user User) error {

	db := database.GetInstance().GetConnection()
	defer db.Close()

	err := db.Save(&user)

	if err != nil {
		return err.Error
	}
	return nil
}

// GetUserByID Return user based on the id provided
func (u *User) GetUserByID(id string) (User, error) {
	db := database.GetInstance().GetConnection()
	defer db.Close()

	user := User{}
	err := db.First(&user, id)

	if err != nil {
		return user, err.Error
	}
	return user, nil
}

// GetUsers Return all user with state active
func (u *User) GetUsers() ([]User, error) {

	db := database.GetInstance().GetConnection()
	defer db.Close()

	var user []User

	err := db.Find(&user)

	if err != nil {
		return user, err.Error
	}
	return user, nil
}

// DeleteUserByID deletes course in database
func (u *User) DeleteUserByID(id string) error {
	db := database.GetInstance().GetConnection()
	defer db.Close()

	err := db.Where("id = ?", id).Delete(Course{})

	if err != nil {
		return err.Error
	}
	return nil
}

// UpdateUser find user and update if exists
func (u *User) UpdateUser(id string, input User) (User, error) {

	db := database.GetInstance().GetConnection()
	defer db.Close()

	user := User{}

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	if err := db.Model(&user).Updates(&input).Error; err != nil {
		return user, err
	}

	return user, nil
}
