package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("file starting")

	content := `ADDing this text in to the file , adding data one more
	added defer to`

	file , err := os.Create("./gomadefile.txt")

	if err != nil {
		// this panic stops the execution of the program and shows the error
		panic(err)
	}

	// now writing the content into the file for that we use io package

	length , err := io.WriteString(file,content)

	if err != nil{
		panic(err)
	}
	fmt.Println("length",length)
	// after this we close the file using 
	defer file.Close()

	readDatafromfileingo("./gomadefile.txt");

}

func readDatafromfileingo(filename string)  {
	
	// reading data from a file we need to use the pacakge io
	datainByte, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}
	// if we print the datainbyte we get the data in numbers to convert it into the real data we use the string()

	fmt.Println(string(datainByte))

}