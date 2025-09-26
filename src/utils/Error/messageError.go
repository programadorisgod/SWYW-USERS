package messageError

import "errors"

var (
	ErrInsertUser       = errors.New("error inserting user into database")
	ErrUserNotFound     = errors.New("user not found")
	ErrSearchingForUser = errors.New("error searching for user, pleasy try again")
)
