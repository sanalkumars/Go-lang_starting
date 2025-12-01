package main

import "fmt"

func main() {
	fmt.Println("welcome to funcations")
	// var num1 int
	// var num2 int
	// fmt.Println("Enter ur first number \t")
	// fmt.Scan(&num1)
	// fmt.Println("Enter ur second number \t")
	// fmt.Scan(&num2)

	// calling the adder funcation for result
	// result :=adder(num1,num2)
	// fmt.Println("result for addition is \t",result)


	// calling  proAdder
	result , message := proAdd(12,23,34,45,23,34)
	fmt.Println("result for slice adding is ",result , message);


}

func adder(num1 , num2  int) int  {
	return num1+num2
}

// if we want to return multiple values of different types 
// instead of return type as int put the types of the return value types in a ()

func proAdd( val ...int) (int , string)  {
	// now the argument is a slice, so for adding we need to loop through
	total :=0

	for _ , v := range val{
		total +=v
	}
	return total , "Hi user , this is ur total"
}