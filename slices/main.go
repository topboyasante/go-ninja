package main

import "fmt"

func main() {

	//Slices are like arrays, but they do not have a fixed length
	var scores = []int{30, 60, 90}
	fmt.Println(scores)

	//Add an element to a slice.
	// Under the hood, this does not change the slice, but creates a new slice based on the old one.
	updated_scores := append(scores, 120)
	fmt.Println(scores, updated_scores)

	//To update a slice, we update it directly:
	scores = updated_scores
	fmt.Println(scores, updated_scores)

	//slice ranges: Use it to get a range of values from a slice:
	top_three := updated_scores[0:3]
	//slice[a:b]: gives you elements from slice[a] to slice[b-1]
	//slice[:b]: gives you elements from slice[0] to slice[b-1]
	//slice[a:]: gives you elements from slice[a] to slice[len(slice)-1]
	fmt.Println(top_three,len(top_three))
}
