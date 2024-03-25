package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	my_name := "asante"

	//These will create new strings based on the provided string
	fmt.Println(strings.Contains("xy", my_name))
	fmt.Println(strings.ReplaceAll(my_name, "xy", "Nana Kwasi"))
	fmt.Println(strings.ToUpper(my_name))

	//gives the index of where the passed string (param 2) is or starts
	fmt.Println(strings.Index(my_name, "a"))

	fmt.Println(strings.Split(my_name, " "))

	ages := []int{30, 160, 180, 120, 150, 90, 60}

	//thos alters the original slice
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 160)
	fmt.Println(index)

	users := []string{"Mario", "Luigi", "Yoshi", "Bowser"}
	sort.Strings(users)
	fmt.Println(users)

	fmt.Println(sort.SearchStrings(users, "Bowser"))

}
