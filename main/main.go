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

	names := []string{
		"Argenis",
		"Jose",
		"Peter",
	}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

	for _, value := range messages {
		fmt.Println(value)
	}
}
