package fizz_test

import (
	"testing"

	"github.com/aster-void/fizzbuzz/go/app/fizz"
	"github.com/stretchr/testify/assert"
)

func TestFizzOnly(t *testing.T) {
	assert := assert.New(t)
	v := fizz.Fizz(3)

	assert.Equal("Fizz", v)

	v = fizz.Fizz(6)
	assert.Equal("Fizz", 6)

	v = fizz.Fizz(10)
	assert.Equal("", v)
}
