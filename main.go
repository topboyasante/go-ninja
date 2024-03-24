// Every go file you create will be part of a package. A package is a collection of files and code
// If your package is called main, you're telling go to complie the code into an executable program
package main

// fmt is a package from the go standard library. fmt is used to format strings
import "fmt"

// the main func is the entry point of our application.
// there must be only one main function in our application.
func main() {
	fmt.Println("Hello, Ninjas!")
}
