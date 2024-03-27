package main

import "fmt"

func main() {

	//This is a map, where the keys are of type "string" and the values are of type "float64"
	menu := map[string]float64{
		"soup":  4.99,
		"pie":   5.99,
		"salad": 3.49,
	}

	fmt.Println(menu)
	fmt.Println(menu["salad"])

	//looping through maps
	for key, value := range menu {
		fmt.Printf("%v : %v \n", key, value)
	}

	// ints as key type
	phonebook := map[int]string{
		230: "Bill",
		490: "Adolf",
		656: "Jersey",
	}

	for key, value := range phonebook {
		fmt.Printf("%v : %v \n", key, value)
	}

	//update an item in a map
	menu["salad"] = 6.45
	fmt.Println(menu)

}
