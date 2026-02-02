package main

import (
	"argeniscuellar/greetings"
	"fmt"
	"log"
)

func main() {
	// Get a greeting message and print it.
	log.SetPrefix("main: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Argenis")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
