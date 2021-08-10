// First Go program
package main

import "fmt"

// Main function
func main() {
	var a [5]int

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
    fmt.Println("length:", len(a))

}
