package main

import "fmt"

func main() {
	x := 0

	for x <= 15 {
		fmt.Println(x)
		x++
	}

	for i := 0; i <= 15; i++ {
		fmt.Println(i)
	}

	users := []string{"Mario", "Luigi", "Yoshi", "Bowser"}
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}

	for index, value := range users {
		fmt.Println(index, value)

	}
}
