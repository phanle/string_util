package string_util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseEmtyStringFine(t *testing.T) {
	assert.Equal(t, Reverse(""), "")
}

func TestReverseSingleCharString(t *testing.T) {
	assert.Equal(t, Reverse("a"), "a")
}

func TestReverseStringWorkForTwoCharStrings(t *testing.T) {
	assert.Equal(t, Reverse("ab"), "ba")
}

func TestReverseStringWorkForThreeCharStrings(t *testing.T) {
	assert.Equal(t, Reverse("abc"), "cba")
}

func TestReverseStringDealsWithUnicodeChars(t *testing.T) {
	assert.Equal(t, Reverse("àộ"), "ộà")
}

func TestReverseWordsDealWithEmptyString(t *testing.T) {
	assert.Equal(t, ReverseWords(""), "")
}

func TestReverseWordsDealsWithOneWordSentence(t *testing.T) {
	assert.Equal(t, ReverseWords("now"), "now")
}

func TestReverseWordsDealsWith2WordStrings(t *testing.T) {
	assert.Equal(t, ReverseWords("blue car"), "car blue")
}
