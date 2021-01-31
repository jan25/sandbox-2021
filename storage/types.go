package main

import "errors"

// ErrFailedToWrite is used when write to a storage failed.
var ErrFailedToWrite = errors.New("write failed")

// ErrRecordNotFound is used when record is not found in storage.
var ErrRecordNotFound = errors.New("record not found")

// Record is a database record.
type Record interface {
	getID() int64
	getValue() string
}

// Writer types are used to write to a concrete storage type.
type Writer interface {
	write(Record) error
}

// Reader types are used to read from a conrete storage type.
type Reader interface {
	read(int64) (Record, error)
}

// Storage refers to a generic storage.
type Storage interface {
	Writer
	Reader
}
