package rand

// Options struct.
type Options struct {
	Reader Reader
}

// Option func.
type Option func(*Options)

// WithReader options to set custom reader.
func WithReader(r Reader) Option {
	return func(o *Options) {
		o.Reader = r
	}
}
