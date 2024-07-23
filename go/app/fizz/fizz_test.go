package fizz_test

import (
	"testing"

	"github.com/aster-void/fizzbuzz/go/app/fizz"
	"github.com/stretchr/testify/assert"
)

func TestFizzOnly(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("Fizz", fizz.Fizz(3))
	assert.Equal("Fizz", fizz.Fizz(6))
	assert.Equal("", fizz.Fizz(10))
}
