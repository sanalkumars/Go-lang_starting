package main

import (
	"fmt"
	"sync"
)

func taskter(id int , wg *sync.WaitGroup)  {
	defer wg.Done()
	fmt.Println("task done is",id)
}

func main() {
	fmt.Println("using wait group in GO")

	// creating a wait group
	var wg sync.WaitGroup

	for i:=0;i<=10;i++{
		// taskter(i) // this is a blocker since this follows linear order
		// solution to avoid this is  go-routines
		wg.Add(1)
		go taskter(i,&wg)
	}

	wg.Wait()
}


