package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	fizzbuzz "github.com/aster-void/fizzbuzz/go/app"
)

func main() {
	var max, min int = 15, 1
	switch len(os.Args) {
	case 1:
		// noop
	case 2:
		var err error
		max, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalln("Failed to parse args:", os.Args[1])
		}
	case 3:
		var err error
		min, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalln("Failed to parse args:", os.Args[1])
		}
		max, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalln("Failed to parse args:", os.Args[2])
		}
	default:
		log.Fatalln("too many args provided. args:", os.Args[1:])
	}

	if max < min {
		log.Fatalln("invalid args: max >= min \nmax =", max, "\nmin =", min)
	}

	list := fizzbuzz.Range(min, max)
	for _, item := range list {
		fmt.Println(item)
	}
}
