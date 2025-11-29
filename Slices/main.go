package main

import "fmt"

func main() {

	// syntax
	//  var slice_name = [] type { values }; we can initialize the slice here , or we can do this later using append method

	// samples

	// var fruits = []string{"apple","orange","grape"}
	// fmt.Printf("type of the slice is %T\n",fruits)
	// fmt.Print("values of the slice is ",fruits,"\n")

	// // adding new element to the slice
	// fruits = append(fruits, "tomato")
	// fmt.Println("updated list ",fruits)

	// // spliting the slice in to a new slice (slice mthod)
	// // fruits = append(fruits[1:]) //skips teh zero th element
	// fmt.Println("new fruit",fruits)

	// newfruits := append(fruits[1:3])
	// fmt.Println("old fruits",fruits) 
	// fmt.Println("new  fruits",newfruits) 


	// creating slice using make()
	scores := make([]int,4)
	fmt.Println(scores)
	scores[0] = 10
	scores[1] = 12
	scores[2] = 134
	scores[3] = 154
	// scores[0] = 15 /*this will result is an error , because when the slice is 
	// being is created we specifically set a size for the slice , when try to assign 
	// value toa memmory which is not yet allocated , will result in error . in these cases the best option is to use the append(), this will allocate
	// new memory and save the value in this space no error*/
	scores = append(scores, 110)

	fmt.Println("updated slice",scores)



}