package db

// Error represents an error condition within the database.
type Error struct {
	message string
	cause   error
}

// Error returns a string representation of the error.
func (e *Error) Error() string {
	if e.cause != nil {
		return e.message + ": " + e.cause.Error()
	}
	return e.message
}
