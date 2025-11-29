package main

import "fmt"

func main() {
	// traditional way
	for i := 0; i < 10; i++ {
		// fmt.Print("[",i ,"]")
	}
	// making for look like while 
	i:=0
	for i<10{
		fmt.Print(i,"\t")
		i++
	}
	// using range 
	s := []int{ 1,2,3,4}
	for i,v := range s {
		fmt.Printf("\n index , value %d %d",i,v)
	}
	// updating slice elements in-place
	for i := range s{
		s[i]=s[i]*i
	}
	fmt.Print("\n",s,"updated slice ")
}