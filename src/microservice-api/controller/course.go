package controller

import (
	"restapi/database"
)

// Course model
type Course struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Status bool   `json:"status,omitempty"`
	Code   string `json:"code,omitempty"`
	UserID int    `json:"userid,omitempty"`
	User   User   `json:"user,omitempty" gorm:"foreignkey:id"`
}

// CreateCourse creates course in database
func (c *Course) CreateCourse(co *Course) error {
	db := database.GetInstance().GetConnection()
	defer db.Close()
	err := db.Save(&co)

	*co = Course{}

	if err != nil {
		return err.Error
	}
	return nil
}

// DeleteCourseByID deletes course in database
func (c *Course) DeleteCourseByID(id string) error {
	db := database.GetInstance().GetConnection()
	defer db.Close()

	err := db.Where("id = ?", id).Delete(Course{})

	if err != nil {
		return err.Error
	}
	return nil
}

// GetCourseByID return a course, based on the id passed by paramtro
func (c *Course) GetCourseByID(id string) (Course, error) {
	db := database.GetInstance().GetConnection()
	defer db.Close()

	course := Course{}

	err := db.First(&course, id)
	db.Model(course).Related(&course.User)

	if err != nil {
		return course, err.Error
	}

	return course, nil
}

// GetAllCourses return all courses that are active
func (c *Course) GetAllCourses() ([]Course, error) {
	db := database.GetInstance().GetConnection()
	defer db.Close()

	courses := []Course{}

	err := db.Preload("User").Find(&courses)

	if err != nil {
		return courses, err.Error
	}
	return courses, nil
}

// UpdateCourse find course anc update especific params
func (c *Course) UpdateCourse(id string, input Course) (Course, error) {
	course := Course{}
	db := database.GetInstance().GetConnection()
	defer db.Close()

	if err := db.Where("id = ?", id).First(&course).Error; err != nil {
		return course, err
	}

	if err := db.Model(&course).Updates(&input).Error; err != nil {
		return course, err
	}

	return course, nil
}
