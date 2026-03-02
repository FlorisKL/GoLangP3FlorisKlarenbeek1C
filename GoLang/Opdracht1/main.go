package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("Welkom bij het nummer gokspel! Raad het nummer tussen 1 en 100.")
}
func gokspel() bool {

	fmt.Println("Hmmm! Welk nummer denk ik nu aan?")

	nummer := rand.Intn(100) + 1
	diff := 1

	fmt.Print("Enter your choice: ")
	fmt.Scan(&diff)

	if diff > 3 || diff < 1 {
		fmt.Println("Ongeldige keuze, defaulting to easy mode.")
		diff = 1
	}
	diffString := "easy"
	chances := 10
	switch diff {
	case 2:
		diffString = "medium"
		chances = 5
	case 3:
		diffString = "hard"
		chances = 3
	}
	fmt.Printf(`Great! You have selected the %s difficulty level. Let's start the game`, diffString)
	for i := 0; i < chances; i++ {
		fmt.Print("Enter your choice: ")
		fmt.Scan(&guess)
		randomNumber(guess, nummer)
		var gok int
		gok = guess

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
}
