package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int

type Number interface {
	~int | int8 | int16 | int32 | int64 | float32 | float64
}

func Min[T Number](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}
func TestMin(t *testing.T) {
	assert.Equal(t, 100, Min[int](200, 100))
	assert.Equal(t, int64(100), Min[int64](int64(200), int64(100)))
	assert.Equal(t, float64(100), Min[float64](float64(200), float64(100)))
	assert.Equal(t, Age(100), Min[Age](Age(200), Age(100)))
}
func TestMinTypeInference(t *testing.T) {
	assert.Equal(t, 100, Min(200, 100))
	assert.Equal(t, int64(100), Min(int64(200), int64(100)))
	assert.Equal(t, float64(100), Min(float64(200), float64(100)))
	assert.Equal(t, Age(100), Min(Age(200), Age(100)))
}
