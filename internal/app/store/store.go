package store

import "errors"

type Store interface {
	User() UserRepository
	School() SchoolRepository
}

var (
	ErrRecordNotFound = errors.New("Record not found")
)
