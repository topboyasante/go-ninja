package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(r *bufio.Reader, prompt string) (s string, e error) {
	fmt.Print(prompt)
	input, error := r.ReadString('\n')
	return strings.TrimSpace(input), error
}

func createBill() bill {
	//Readers are used to get user input through the terminal.
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput(reader, "Create a New Bill Name:")

	b := generateNewBill(name)
	fmt.Printf("Created bill %v \n", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput(reader, "Choose an option: \n A - Add Item \n S - Save Bill \n T - Add Tip \n")
	fmt.Println(opt)

	switch opt {
	case "a":
		name, _ := getInput(reader, "What is the item name? \n")
		price, _ := getInput(reader, "How much is it? \n")

		//convert strings to float
		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("Invalid Price value")
			promptOptions(b)
		}

		b.addItem(name, p)

		fmt.Println("Item added!")

		promptOptions(b)
	case "t":
		tip, _ := getInput(reader, "How much do you want to tip?")

		p, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("Invalid Price value")
			promptOptions(b)
		}

		b.updateTip(p)

		fmt.Println("Tip Updated!")

		promptOptions(b)
	case "s":
		b.saveBill()
	default:
		fmt.Println("Invalid Input")
		promptOptions(b) //looks like recursion
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
	fmt.Println(mybill)
}
