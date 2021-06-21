package rand

// Reader interface.
type Reader interface {
	Read(p []byte) (n int, err error)
}
