package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("welcome to user inputs")

	// using fmt.scan()  this is for getting a single value from the user

	// var n int 
	// fmt.Println("enter the number ")
	// fmt.Scan(&n)
	// fmt.Println("entered number is ",n)

	// for reading an entire line  use scanln from fmt this exit when the first white space occure in the below case the output will be only "i"
	var line string
	fmt.Println("entire ur text")
	// fmt.Scanln(&line)
	
	// for getting the entire line with out an exit on space use bufio.scanner
	scanner:=bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line = scanner.Text()
	
	fmt.Println("entered text is ",line)

}