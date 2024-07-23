package app_test

import (
	"testing"

	fizzbuzz "github.com/aster-void/fizzbuzz/go/app"
	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("1", fizzbuzz.Apply(1))
	assert.Equal("Fizz", fizzbuzz.Apply(3))
	assert.Equal("Buzz", fizzbuzz.Apply(5))
	assert.Equal("FizzBuzz", fizzbuzz.Apply(15))
}
