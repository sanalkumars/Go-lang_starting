package main

import "fmt"

func main() {

	const age = 20 //cannot be re-initialised

	// constant grouping

	const (
		host int    = 3000
		url  string = " http ://localhost"
	)

	fmt.Printf("host is %d\n",host)
	fmt.Printf("url is %s",url)

	// creating dynamic strings

	var port = 3000
	const host2 = "http://localhost:"

	// fmt.Printf("%s",host2,"%d",port) this way of doing is wrong
	fmt.Printf("\n%s%d\n",host2,port)


}