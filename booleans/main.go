package main

import "fmt"

func main() {
	age := 50

	fmt.Println(age > 50)

	if age > 30 {
		fmt.Println("You are old. Ew")
	} else if age == 40 {
		fmt.Println("40? EWWW")
	} else {
		fmt.Println("You're safe. I guess")
	}

	users := []string{"Mario", "Luigi", "Yoshi", "Bowser"}
	for index, value := range users {
		if index == 1 {
			{
				fmt.Println("Starting from index 1")
				continue
			}
		}
		if index > 3 {
			fmt.Println("Starting from index 1")
			break
		}

		fmt.Println(value)

	}

}
