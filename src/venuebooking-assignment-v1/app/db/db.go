package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type ReaderDB struct {
	db *sql.DB
}

type WriterDB struct {
	*ReaderDB
}

var (
	reader *ReaderDB
	writer *WriterDB
)

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
