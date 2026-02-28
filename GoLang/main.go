package main

import (
	"fmt"
)


func main() {

fmt.Println("Hmmm! Welk nummer denk ik nu aan?")

nummer := rand.Intn(100) + 1

var gok int

	if gok == nummer {
		fmt.Println("Gefeliciteerd! Je hebt het juiste nummer geraden!")
	} if gok < nummer {
		fmt.Println("Helaas, het nummer is hoger. Probeer het opnieuw.")

	}

}