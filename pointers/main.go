package main

import "fmt"

func updateName(x *string) {

	//*x is the value the string pointer is pointing to
	*x = "Wedge"
}

func main() {
	name := "Asante"


	fmt.Println("Memory Address of name:", &name)

	//stores the memory address of the name variable
	m := &name
	fmt.Println("Memory Address of m:", m)

	//Get the value the pointer points to
	fmt.Println("Value at Memory Address of m:", *m)

	updateName(m)

	fmt.Println(name)

}
