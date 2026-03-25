package main

import (
	"fmt"
	"strings"
)

func main1() {
	fmt.Println("Hello world")
}

naam1 := "John"
leeftijd1 := 30
nummer1 := "SD-001"

naam2 := "Jane"
leeftijd2 := 25
nummer2 := "SD-002"

type Student struct {
	Naam     string
	Leeftijd int
	Nummer   string
}

S := Student{
	Naam:     "Alice",
	Leeftijd: 19,
}


// GetEvenNumbers verzamelt alle even getallen van 1 tot en met limit
func GetEvenNumbers(limit int) []int {
	evens := []int{} 

	// TODO: Schrijf een for-loop die van 1 tot en met limit loopt
	for i := 1; i <= limit; i++{
		if i%2 == 0 {
			evens = append(evens, i)
		}
	}
	// TODO: Gebruik een if-statement en de modulo operator (%) om te checken of het getal even is
	// TODO: Voeg het getal toe aan de 'evens' slice met append(evens, getal)

	return evens
}

func main2() {
	// Roep de functie aan met 20 en print het resultaat
	result := GetEvenNumbers(20)
	fmt.Println("Even getallen:", result)
}







func main3() {
	full, length := TransformName("  Jan  ", "  Modaal")
	fmt.Printf("Naam: %q, Totale lengte: %d\n", full, length)
}

func TransformName(firstName string, lastName string) (string, int) {
	f := strings.TrimSpace(firstName)
	l := strings.TrimSpace(lastName)
	return f + " " + l, len(f + l)
}

// TODO: Schrijf de functie Optellen
func Optellen(a, b int) int {
	return a + b
}
// TODO: Schrijf de functie Aftrekken
func Aftrekken(a, b int) int {
	return a - b
}
// TODO: Schrijf de functie Vermenigvuldigen
func Vermenigvuldigen(a, b int) int {
	return a * b
}
// TODO: Schrijf de functie Delen (let op de float64 conversie!)
func Delen(a, b int) float64 {
	return float64(a) / float64(b)
}
// TODO: Schrijf de functie VerdeelEnHeers die twee ints teruggeeft
func VerdeelEnHeers(a, b int) (int, int){
	q := a / b
	r := a % b
	return q,r
}
func main4() {
	fmt.Println("Optellen 5 + 3:", Optellen(5, 3))
	fmt.Println("Aftrekken 10 - 4:", Aftrekken(10, 4))
	fmt.Println("Vermenigvuldigen 6 * 7:", Vermenigvuldigen(6, 7))
	fmt.Println("Delen 10 / 4:", Delen(10, 4))

	q, r := VerdeelEnHeers(10, 3)
	fmt.Printf("10 gedeeld door 3 is %d met rest %d\n", q, r)
}




// We definiëren de 'blauwdruk' voor een GameCharacter
type GameCharacter struct {
    Name   string
    Health int
    Level  int
    IsHero bool
}

func main5() {
    // Hier maken we een 'instance' van onze struct
    player1 := GameCharacter{
        Name:   "Gopher de Grote",
        Health: 100,
        Level:  5,
        IsHero: true,
    }

    fmt.Println("Naam:", player1.Name)
    fmt.Println("Levenspunten:", player1.Health)
}







type Engine struct {
    Horsepower int
    Type       string
}

type Car struct {
    Brand  string
    Model  string
    Engine Engine // Een struct binnen een struct!
}

func main6() {
    myCar := Car{
        Brand: "Tesla",
        Model: "Model 3",
        Engine: Engine{
            Horsepower: 300,
            Type:       "Electric",
        },
    }

    fmt.Printf("Mijn %s heeft %d pk!\n", myCar.Brand, myCar.Engine.Horsepower)
}


// TODO: Definieer de User struct met velden Name, Age en Email
type User struct {
	Name string
	Age int
	Email string
}
// TODO: Maak de CreateUser functie die een User retourneert
func CreateUser(Name string, Age int, Email string) User {
	return User{
		Name: Name,
		Age: Age,
		Email: Email,
	}
}
// TODO: Maak de PrintUser functie die de data print in het juiste formaat
func PrintUser(u User) {
	fmt.Printf("Naam: %s, Leeftijd: %d, Email: %s/n", u.Name, u.Age, u.Email)
}
func main() {
	// TODO: Maak een nieuwe gebruiker aan met CreateUser
	user := CreateUser("Jan", 30, "Jan@email.com")
	// TODO: Print de gebruiker met PrintUser
	PrintUser(user)
}