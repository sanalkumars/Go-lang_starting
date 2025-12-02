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
	Status string
}

func main() {
	fmt.Print("structure in go lang \n")

  // Method 1: Named fields (recommended)
    sanal := User{
        Name:   "Sanal",
        Age:    23,
        Address: Address{City: "halo", House_No: 212},
        Status: "active",
    }
	sanal.getUserStatus()
	sanal.changeStatus()
	sanal.changeStatusPermenent()
	sanal.getUserStatus()

	 // Method 2: Positional values
    john := User{"John", 25, Address{101,"New York" }, "active"}

    // Method 3: Using pointer to create object
    mike := &User{
        Name:   "Mike",
        Age:    30,
        Status: "inactive",
    }

    fmt.Println(sanal)
    fmt.Println(john)
    fmt.Println(mike)


}

func (u User)getUserStatus()  {
	fmt.Println("user status is ",u.Status)
}

// funcation for manupulating the values for the created object/structure 

// this is tempporary as we only pass a copy of the original structure
func (u User)changeStatus()  {
	u.Status ="inactive" // this change is only temporary, as we are only passing the copy of the original object
	fmt.Println("new status is ",u.Status) 
}
// here instead of the copy we are passing a pointer to the original object
func (u *User)changeStatusPermenent()  {
	u.Status ="pending" // this change is only temporary, as we are only passing the copy of the original object
	fmt.Println("new status is ",u.Status) 
}
