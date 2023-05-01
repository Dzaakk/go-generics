package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func ExperimentalMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestExperimentalMin(t *testing.T) {
	assert.Equal(t, 100, ExperimentalMin(100, 200))
	assert.Equal(t, 100.00, ExperimentalMin(100.00, 200.00))
}

func TestExperimentalMaps(t *testing.T) {

	first := map[string]string{
		"name": "User1",
	}

	second := map[string]string{
		"name": "User1",
	}

	assert.True(t, maps.Equal(first, second))
}

func TestExperimentalSlice(t *testing.T) {
	first := []string{"User1"}
	second := []string{"User1"}

	assert.True(t, slices.Equal(first, second))
}
