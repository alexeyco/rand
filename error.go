package rand

// Error structure.
type Error struct {
	message string
}

// Error returns error message.
func (e Error) Error() string {
	return e.message
}
