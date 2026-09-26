package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	log.SetPrefix("greetings:   ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	fmt.Println(quote.Go())
}
