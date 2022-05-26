package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set the properties if the predefined Logger, including
	// the log emtry prefix and a flag to disable printing
	// the time, source file, and line number

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// request a greeting message
	message, err := greetings.Hello("")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print it to the console
	fmt.Println(message)
}
