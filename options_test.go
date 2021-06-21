package rand_test

import (
	gorand "crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alexeyco/rand"
)

func TestWithReader(t *testing.T) {
	t.Parallel()

	var o rand.Options

	rand.WithReader(gorand.Reader)(&o)

	assert.NotNil(t, o.Reader)
	assert.Equal(t, gorand.Reader, o.Reader)
}
