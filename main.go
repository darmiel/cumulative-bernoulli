package main

import (
	"log"
)

func main() {
	res := findUpperBoundLe(1200, 0.2, 0.025)
	if res != nil {
		log.Println("Result:", *res)
	} else {
		log.Println("Nothing found :(")
	}

	res = findUpperBoundGe(1200, 0.2, 0.025)
	if res != nil {
		log.Println("Result:", *res)
	} else {
		log.Println("Nothing found :(")
	}
}
