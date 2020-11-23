CREATE database my_db;

USE my_db;
CREATE TABLE Users (ID varchar (5) NOT NULL PRIMARY KEY, FirstName VARCHAR(30), LastName VARCHAR(30), Age INT, Api_key VARCHAR(100));
INSERT INTO Users (ID, FirstName, LastName, Age, Api_key) VALUES ("0001", "WeiMeng", "Lee", 25, "10239qwede2339121821dasflwknw" );
SELECT * FROM Users;

CREATE TABLE Course (ID varchar(6) NOT NULL PRIMARY KEY, CourseName VARCHAR(200) KEY, CourseStatus BOOLEAN NOT NULL, CourseDay INT(2), CourseMonth INT(2), CourseYear INT(4), CONSTRAINT UsersBooking FOREIGN KEY CourseName REFERENCES Users(ID));
INSERT INTO Course (ID, CourseName, CourseStatus, CourseDay, CourseMonth, CourseYear) VALUES ("IOS101","IOS Programming 101", True, 30, 11, 2020);



