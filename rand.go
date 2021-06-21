// Package rand provides cryptographically secure random numbers and strings generator.
package rand

import (
	"crypto/rand"
)

// New returns new provider instance.
func New(options ...Option) *Random {
	o := Options{
		Reader: rand.Reader,
	}

	for _, option := range options {
		option(&o)
	}

	return &Random{
		reader: o.Reader,
	}
}
