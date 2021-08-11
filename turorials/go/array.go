// First Go program
package main

import "fmt"

// Main function
func main() {
	// Init an int array, size 5
	var a[5] int

	a[4] = 100
	fmt.Println("===\nset:", a)
	fmt.Println("get item [4]:", a[4])
    fmt.Println("length:", len(a))

    // Declare and init array in one line - pretty cool
	b := [5] int{1,2,3,4,5}
	fmt.Println("====\ndcl:", b)

	// Array types are one dimensional but can be used 
	// to compose multi-dimensional data structures
	var twoDArray [2][3] int
	for i :=0; i<2; i++{
		for j:=0;j<3;j++{
			twoDArray [i][j]= i+j
		}
	}
	
	fmt.Println("2d", twoDArray)
	
}
