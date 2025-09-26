package messageError

import "errors"

var (
	ErrorHashingPassord = errors.New("Error hashing password")
	ErrInsertUser       = errors.New("error inserting user into database")
	ErrUserNotFound     = errors.New("user not found")
	ErrSearchingForUser = errors.New("error searching for user, pleasy try again")
)
