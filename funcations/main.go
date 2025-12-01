package main

import "fmt"

func main() {
	fmt.Println("welcome to funcations")
	var num1 int
	var num2 int
	fmt.Println("Enter ur first number \t")
	fmt.Scan(&num1)
	fmt.Println("Enter ur second number \t")
	fmt.Scan(&num2)

	// calling the adder funcation for result
	result :=adder(num1,num2)
	fmt.Println("result for addition is \t",result)

}

func adder(num1 , num2  int) int  {
	return num1+num2
}