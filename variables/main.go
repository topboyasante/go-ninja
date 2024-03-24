package main

import "fmt"

func main() {

	//Variables
	var user string = "Nana Kwasi Asante"
	var user_2 = "Hannibal Lecter"
	var user_3 string

	//You can't declare variables like this outside of a function
	user_4 := "Vector"

	fmt.Println(user, user_2, user_3, user_4)

	user_3 = "Jared Leto"

	var age_1 int = 30
	age_2 := 50
	var age_3 = 60

	fmt.Println(age_1, age_2, age_3)

	var num_1 int8 = -25
	var num_2 int16 = 300
	var num_3 uint8 = 100 // no negative values allowed for uints

	fmt.Println(num_1, num_2, num_3)

	var score_1 float32 = 40
	var score_2 float64 = 4.0
	score_3 := 5.0

	fmt.Println(score_1, score_2, score_3)
}
