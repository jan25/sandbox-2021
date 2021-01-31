package types

import "errors"

// ErrFailedToWrite is used when Write to a storage failed.
var ErrFailedToWrite = errors.New("Write failed")

// ErrRecordNotFound is used when record is not found in storage.
var ErrRecordNotFound = errors.New("record not found")

// Record is a database record.
type Record interface {
	GetID() int64
	GetValue() string
}

// Writer types are used to Write to a concrete storage type.
type Writer interface {
	Write(Record) error
}

// Reader types are used to Read from a conrete storage type.
type Reader interface {
	Read(int64) (Record, error)
}

// Storage refers to a generic storage.
type Storage interface {
	Writer
	Reader
}
