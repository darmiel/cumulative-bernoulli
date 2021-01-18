package main

import (
	"log"
)

func main() {
	log.Println(F(2000, 0.2, 300))

	res := findUpperBoundLe(2000, 0.2, 0.025)
	if res != nil {
		log.Println("Result:", *res)
	} else {
		log.Println("Nothing found :(")
	}
}
