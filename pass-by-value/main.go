package main

import "fmt"

func updateName(x string) {
	x = "Wedge"
}

func updateMenu(menu map[string]float64) {
	menu["soup"] = 3.69
}

func main() {
	//GROUP A TYPES: When you pass a group A variable into a function as a parameter, go creates a copy of that variable and works with that, instead of the actual variable.
	name := "Asante"

	updateName(name)

	fmt.Println(name) //returns "Asante"

	//GROUP A TYPES:
	// When you pass a group B variable into a function as a parameter, go crates a pointer to that variable, and stores both the pointer and variable in memory.
	// When you pass such variables into a function, go creates a copy of the pointer, uses the copy to "point" to the variable, and then updates it.
	menu := map[string]float64{
		"soup":  4.99,
		"pie":   5.99,
		"salad": 3.49,
	}

	updateMenu(menu)

	fmt.Println(menu)
}
