package main

import (
	"fmt"
	"harshrpg/greetings"
	"log"
)

func main() {
	// fmt.Println(quote.Hello())
	// fmt.Println(quote.Go())
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Harsh", "John", "Jack"}
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
