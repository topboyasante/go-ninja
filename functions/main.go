package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	sayGreeting("Asante")
	sayBye("Asante")
	users := []string{"Mario", "Luigi", "Yoshi", "Bowser"}

	cycleName(users, sayGreeting)

	firstNameInitial, secondNameInitial := getInitials("Nana Kwasi")
	fmt.Println(firstNameInitial, secondNameInitial)

	a1 := circleArea(3.5)
	fmt.Println(a1)
}

func sayGreeting(name string) {
	fmt.Printf("Good Morning, %v \n", name)
}

func sayBye(name string) {
	fmt.Printf("Good Bye, %v \n", name)
}

func cycleName(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

// Functions returning multple values
func getInitials(n string) (string, string) {
	namesToUpperCase := strings.ToUpper(n)
	names := strings.Split(namesToUpperCase, " ")

	var initials []string

	for _, values := range names {
		initials = append(initials, values[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}
