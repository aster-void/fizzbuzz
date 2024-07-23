package buzz_test

import (
	"testing"

	"github.com/aster-void/fizzbuzz/go/app/buzz"
	"github.com/stretchr/testify/assert"
)

func TestBuzz(t *testing.T) {
	assert := assert.New(t)

	v := buzz.Buzz(5)
	assert.Equal("Buzz", v)

	v = buzz.Buzz(10)
	assert.Equal("Buzz", v)

	v = buzz.Buzz(12)
	assert.Equal("", v)
}
