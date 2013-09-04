package string_util

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMinimumDistanceDealsWithEmptyStrings(t *testing.T) {
	assert.Equal(t, MinimumDistance("", ""), 0)
}

func TestMinimumDistanceDealsWithSingleCharStringAndEmptyString(t *testing.T) {
	assert.Equal(t, MinimumDistance("a", ""), 1)
}

func TestMinimumDistanceDealsWithEmptyStringAndSingleCharString(t *testing.T) {
	assert.Equal(t, MinimumDistance("", "a"), 1)
}

func TestMinimumDistanceDealsWithTwoSingleCharStrings(t *testing.T) {
	assert.Equal(t, MinimumDistance("a", "b"), 1)
}

func TestMinimumDistanceDealsAMultipleStringAndASingleCharString(t *testing.T) {
	assert.Equal(t, MinimumDistance("ca", "c"), 1)
	assert.Equal(t, MinimumDistance("cat", "c"), 2)
	assert.Equal(t, MinimumDistance("chat", "cat"), 1)
}


