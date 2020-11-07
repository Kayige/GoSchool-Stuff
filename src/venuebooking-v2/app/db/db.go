// package db provides new connection to database
// uses *ReaderDB for read(SELECT) only queries
// uses *WriterDB for write(INSERT,UPDATE,DELETE) only queries

package db

import (
	"database/sql"
	"log"
	"time"

	// imports mysql for gos
	_ "github.com/go-sql-driver/mysql"
)

// ReaderDB struct having db instance
type ReaderDB struct {
	db *sql.DB
}

// WriterDB struct has an embeded type having db instance
type WriterDB struct {
	*ReaderDB
}

var (
	reader *ReaderDB
	writer *WriterDB
)

// Reader makes and return new db connection
// should be used for read(SELECT) queries
func Reader(connectionStr string) (*ReaderDB, error) {
	if reader != nil {
		return reader, nil
	}

	db, err := new(connectionStr)
	if err != nil {
		return nil, err
	}

	return &ReaderDB{db: db}, nil
}

// Writer makes and return new db connection
// should be used for write(INSERT,UPDATE,DELTE) queries
func Writer(connectionStr string) (*WriterDB, error) {
	if writer != nil {
		return writer, nil
	}

	db, err := new(connectionStr)
	if err != nil {
		return nil, err
	}

	return &WriterDB{&ReaderDB{db: db}}, nil
}

// new takes connection url as a string
// make new db connection and return new db instance
func new(connectionStr string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connectionStr)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err = db.Ping(); err != nil {
		log.Fatalf("unable to connect to database, error : %v", err)
	}
	return db, nil
}
