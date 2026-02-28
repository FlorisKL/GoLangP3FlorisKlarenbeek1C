package main

import (
	"fmt"
)


func main() {

fmt.Println("Welkom bij het nummer gokspel! Raad het nummer tussen 1 en 100.")
} 
func gokspel() bool {

	fmt.Println("Hmmm! Welk nummer denk ik nu aan?")

nummer := rand.Intn(100) + 1

	var gok int
if gok < nummer {
		fmt.Println("Helaas, het nummer is hoger. Probeer het opnieuw.")
		return false

	} else if gok > nummer {
		fmt.Println("Helaas, het nummer is lager. Probeer het opnieuw.")
		return false

	} else if gok == nummer {
		fmt.Println("Gefeliciteerd! Je hebt het juiste nummer geraden!")
		return true
	}
	return false
}