package rand_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/alexeyco/rand"
)

func TestRandom_Int64(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	readerMock := NewMockReader(ctrl)

	expectedError := errors.New("error")

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		n, err := rand.New().Int64(10)

		assert.GreaterOrEqual(t, n, int64(0))
		assert.Less(t, n, int64(10))
		assert.NoError(t, err)
	})

	t.Run("MaxIsEqualToMin", func(t *testing.T) {
		t.Parallel()

		n, err := rand.New().Int64(10, 10)

		assert.Equal(t, int64(10), n)
		assert.NoError(t, err)
	})

	t.Run("MaxLessThanMin", func(t *testing.T) {
		t.Parallel()

		n, err := rand.New().Int64(10, 100)

		assert.Empty(t, n)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &rand.Error{})
		assert.EqualError(t, err, "max should be greater than min")
	})

	t.Run("ReaderError", func(t *testing.T) {
		t.Parallel()

		readerMock.EXPECT().Read(gomock.Any()).Return(0, expectedError)
		n, err := rand.New(rand.WithReader(readerMock)).Int64(10)

		assert.Empty(t, n)
		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
	})
}

func TestRandom_String(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	readerMock := NewMockReader(ctrl)

	expectedError := errors.New("error")

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		s, err := rand.New().String(rand.LettersAndNumbers, 10, 10)

		t.Log(s)
		assert.Len(t, s, 10)
		assert.NoError(t, err)
	})

	t.Run("EmptyAlphabet", func(t *testing.T) {
		t.Parallel()

		s, err := rand.New().String("", 10)

		assert.Empty(t, s)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &rand.Error{})
		assert.EqualError(t, err, "alphabet shouldn't be empty string")
	})

	t.Run("ReaderError", func(t *testing.T) {
		t.Parallel()

		readerMock.EXPECT().Read(gomock.Any()).Return(0, expectedError)
		s, err := rand.New(rand.WithReader(readerMock)).String(rand.LettersAndNumbers, 10)

		assert.Empty(t, s)
		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
	})

	t.Run("ReaderError2", func(t *testing.T) {
		t.Parallel()

		readerMock.EXPECT().Read(gomock.Any()).Return(0, expectedError)
		s, err := rand.New(rand.WithReader(readerMock)).String(rand.LettersAndNumbers, 10, 10)

		assert.Empty(t, s)
		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
	})
}
