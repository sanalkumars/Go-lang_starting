package main

import "fmt"

// in-case we have a string slice we need to create a new funcation, since we can't use the below so we use generics
// func printslice( items []int)  {
	
// 	for _,item := range items {
// 		fmt.Println(item)
// 	}
// }below funcations uses generics

// 
func printslice[T any]( items []T)  {
	
	for _,item := range items {
		fmt.Print(item,"\t")
	}
	fmt.Println()
}
func main() {
	fmt.Println("welcome to generics in go")

	printslice([]int{1,2,3})
	printslice([]string{"2","1","3"})
}