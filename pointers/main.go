package main

import "fmt"

func main() {
	fmt.Println("basics of pointers");

	var ptr *int
	var sum int =10
	ptr = &sum
	*ptr = *ptr+2
	fmt.Println("memmory for sum is ",ptr)
	fmt.Println("memmory for sum is ",*ptr)
	printPointerValue(ptr)
}
func printPointerValue(ptr *int) {
    fmt.Println("Value from function:", *ptr)
}
