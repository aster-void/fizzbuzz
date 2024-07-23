package main

import (
	"fmt"
	fizzbuzz "github.com/aster-void/fizzbuzz/go/app"
)

func main() {
	list := fizzbuzz.Range(1, 15)
	for _, item := range list {
		fmt.Println(item)
	}
}
