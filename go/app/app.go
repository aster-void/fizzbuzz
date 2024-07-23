package app

import (
	"strconv"

	"github.com/aster-void/fizzbuzz/go/app/buzz"
	"github.com/aster-void/fizzbuzz/go/app/fizz"
)

func Range(start int, last int) []string {
	length := last - start + 1
	src := make([]int, length)
	for i := 0; i < length; i++ {
		src[i] = i + start
	}
	return Map(src)
}

func Map(src []int) []string {
	length := len(src)
	dest := make([]string, length)
	for i := 0; i < length; i++ {
		dest[i] = Apply(src[i])
	}
	return dest
}

func Apply(i int) string {
	s := fizz.Fizz(i) + buzz.Buzz(i)
	if s == "" {
		return strconv.Itoa(i)
	} else {
		return s
	}
}
