package utilities

import "errors"

var (
	// ErrorID when id is not valid
	ErrorID = errors.New("Id not valid")

	// ErrorUserID when user id in context is not valid
	ErrorUserID = errors.New("User Unauthorized")
)
