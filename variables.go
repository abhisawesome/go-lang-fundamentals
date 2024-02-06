// Declare variables of different types (int, float64, string, bool) and print their values.

package main

import "fmt"

func main(){
	var i int
	var f float64
	var b bool
	var s string
	
	i=10
	f=20
	b=true
	s="Hello"

	newValue := 100
	
	
	fmt.Println("int value is:",i)
	fmt.Println("float value is:",f)
	fmt.Println("bool value is:",b)
	fmt.Println("string value is:",s)

	fmt.Println("new value is:",newValue)

}