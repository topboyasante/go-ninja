package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make a new bill:
func generateNewBill(name string) bill {
	new_bill := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "eggs": 5.99},
		tip:   0,
	}

	return new_bill
}

// reciever functions: simple
func (b bill) format() string {
	header := "Bill Breakdown \n" 
	var total float64

	for key, value := range b.items {
		header += fmt.Sprintf("%-25v ...$%v \n", key+":", value)
		total += value
	}

	header += fmt.Sprintf("%-25v ...$%0.2f", "Total:", total)
	return header
}

//%-25v make that variable 25 characters long to the right. make it positive and the spaces are on the left
