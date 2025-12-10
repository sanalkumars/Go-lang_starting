package main

import (
	"fmt"
	"time"
)

func taskter(id int)  {
	fmt.Println("task done is",id)
}

func main() {
	fmt.Println("welcome to go routines")

	for i:=0;i<=10;i++{
		// taskter(i) // this is a blocker since this follows linear order
		// solution to avoid this is  go-routines
		go taskter(i)
	}

	// since now each funcation call, main function runs in seperate thread we wont get the result unless we make the main funcation wait

	time.Sleep(time.Second *2)
}