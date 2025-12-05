package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/users"

func main() {
	fmt.Println("first web request")

	response , err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println("reponse for from is ",response)
	fmt.Printf("reponse type %T ",response)

	// by using defer it make sure this runs/executes after all the other code execution
	defer response.Body.Close() // this is the reponsibility of the developer to close the connection after reciving the response , wont disconnect automatically

	data , err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("data from url :",string(data))
}