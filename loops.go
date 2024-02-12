// Create simple loops (for, while) to iterate over a range of values or repeat a block of code.


package main

import "fmt"

func main() {
	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}
}
