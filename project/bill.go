package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make a new bill:
func generateNewBill(name string) bill {
	new_bill := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return new_bill
}

// reciever functions: simple

// Get the breakdown of your bill
func (b *bill) format() string {
	header := "Bill Breakdown \n"
	var total float64

	for key, value := range b.items {
		header += fmt.Sprintf("%-25v ...$%v \n", key+":", value)
		total += value
	}

	header += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", b.tip)

	header += fmt.Sprintf("%-25v ...$%0.2f", "Total:", total+b.tip)
	return header
}

//%-25v make that variable 25 characters long to the right. make it positive and the spaces are on the left

func (b *bill) updateTip(tip float64) {
	b.tip = tip
	// or you can do this: (*b).tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

//for structs pointers are automatically dereferenced

func (b *bill) saveBill() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		{
			panic(err)
		}
	}

	fmt.Println("bill was saved.")
}
