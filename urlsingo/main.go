package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://jsonplaceholder.typicode.com/todos/1?username=abcdefg&age=21"

func main() {
	fmt.Println("manipulating url in go")
	fmt.Println("url is ",myurl)

	// parsing data from the url /extracting data from url

	data, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	// fmt.Println("url parsed info is ",data.Scheme)
	// fmt.Println("url parsed info is ",data.Host)
	// fmt.Println("url parsed info is ",data.Path)
	// fmt.Println("url parsed info is ",data.Hostname())
	// fmt.Println("url parsed info is ",data.Port())
	// fmt.Println("url parsed info is ",data.RawQuery)

	qparams :=data.Query()
	fmt.Println("query",qparams)
	// for accessing each value 
	
	fmt.Println("query",qparams["age"])
	fmt.Println("query",qparams["username"])

}