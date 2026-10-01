package main

import (
	"fmt"
	"log"
	"example/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and print it.
	message, err := greetings.Hello("")
	//If error returned, print it and exit
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
