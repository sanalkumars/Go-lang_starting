package main

import "fmt"

type Address struct{
	House_No int
	City string
}

type User struct {
	Name    string
	Age     int
	Address Address
}

func main() {
	fmt.Print("structure in go lang \n")
	var student User
	student.Name ="Sanal"
	 student.Age = 25
    
	//  creating new structure in one line 

	student1 := User{ Name: "sanal",Age: 23,Address: Address{House_No: 102,City: "Ala"}}

	fmt.Println("student created using struct User",student)
	fmt.Println("student created using struct User + nested structure",student1)

	// accessing nested values
	// fmt.Println("nested values",student1.Address.City,student1.Address.House_No)

	// methods  on structure 
	// calling the function
	student1.printStructInfo()

}

func (u User) printStructInfo(){
	fmt.Println("accessing structure inside a funcation got by argument",
	u.Name )
}