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
	scores := make([]int, 4)
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

	// fmt.Println("updated slice",scores)

	// duplicating a slice
	dup := make([]int, len(scores)) //creates a new slice with the same length as the source we given, but no value

	// to copy the value into the new slice use copy()
	copy(dup, scores)
	// fmt.Println("duplicated slice",dup)

	// looping through the slice using for-range
	for i, val := range scores {
		fmt.Printf("index = %d , value is %d \n", i, val)
	}

	// clearing/reseting the created slice to empty
	dup = dup[:0]
	// fmt.Println("cleared duplicate slice is ",dup)

	// slice creatde using new method , this is not a recommented method
	// newslice := new([]int)
	// // fmt.Print("new based slice",newslice)
	// // to assing value to the slice created using new need to use make to give a size then assign value
	// *newslice = make([]int, 3)

	// (*newslice)[0] =111
	// fmt.Print(newslice)

	// remove a value from a slice based on index
	test := make([]int,0)
	test = append(test, 1, 2, 3, 4, 5)
	// element to remove from the list is in index 3 (value -4 )
	i :=3
	fmt.Print(test,"test before remove")
	test = append(test[:i],test[i+1:]...)


	fmt.Print(test," test after remove")
}
