package main

import "fmt"

func main() {
	myBill := generateNewBill("Hitler's Bill")

	fmt.Println(myBill.format())

}
