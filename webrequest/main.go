package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

	// data , err := ioutil.ReadAll(response.Body)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("data from url :",string(data))

	// another way of getting the data from the response is 

	var responseStr strings.Builder

	content, _ := ioutil.ReadAll(response.Body)
	count,_:= responseStr.Write(content)
	fmt.Print(count)
	fmt.Println(responseStr.String())
}