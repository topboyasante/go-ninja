package main

import "fmt"

func main() {

	user_3 := "Asante"

	//fmt package:
	fmt.Print("Hello, ")
	fmt.Print("World! \n")
	fmt.Print("Lorem")

	fmt.Println("Hello, Asante")
	fmt.Println("Goodbye, Asante")

	age_7 := 50
	fmt.Println("My age is ", age_7)

	// %_ = format specifier
	// %v - variables
	// %T - data type
	// %q - strings
	// %xf - float, where x is a decimal that tells the number of decimal places we need. 0.2 is 2DP
	fmt.Printf("My name is %v", user_3)

	// Sprintf saves formatted strings
	var str = fmt.Sprint("My name is %v", user_3)
	fmt.Println(str)
}
