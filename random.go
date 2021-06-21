package rand

import (
	"crypto/rand"
	"math/big"

	"github.com/pkg/errors"
)

// Random provider.
type Random struct {
	reader Reader
}

// Int64 returns random int64.
func (r *Random) Int64(max int64, min ...int64) (int64, error) {
	var from int64
	if len(min) > 0 {
		from = min[0]
	}

	if from == max {
		return max, nil
	}

	if from > max {
		return 0, errors.WithStack(Error{
			"max should be greater than min",
		})
	}

	n, err := rand.Int(r.reader, big.NewInt(max-from))
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return from + n.Int64(), nil
}

// String returns random stwring.
func (r *Random) String(alphabet Alphabet, max int64, min ...int64) (string, error) {
	if alphabet == "" {
		return "", errors.WithStack(Error{
			"alphabet shouldn't be empty string",
		})
	}

	length, err := r.Int64(max, min...)
	if err != nil {
		return "", errors.WithStack(err)
	}

	runes := []rune(alphabet)

	l := len(runes)
	res := make([]rune, length)

	for i := int64(0); i < length; i++ {
		n, err := r.Int64(int64(l))
		if err != nil {
			return "", err
		}

		res[i] = runes[int(n)]
	}

	return string(res), nil
}
