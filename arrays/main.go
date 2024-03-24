package main

import "fmt"

func main() {

	//An array has a fixed length and type. it cannot be changed once declared.
	var ages [3]int = [3]int{30, 60, 90}
	var net_worth = [3]int{40000, 809000, 1256512}
	users := [4]string{"Mario", "Luigi", "Yoshi", "Bowser"}

	//	Get the length of an array using the len() function:
	fmt.Println(ages, len(ages))
	fmt.Println(net_worth, len(net_worth))
	fmt.Println(users, len(users))

	//Change array values
	users[1] = "Princess Peach"

}
