package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Password string   `json:"-"`      // "-" hides this field in JSON
	Content  []string `json:"content"`
}

func main() {
	fmt.Println("welcome to json in GO ")
	encodeJSON()
}

// encoding json ( converting Go data to JSON )
func encodeJSON() {

	courses := []Course{
		{
			Id:      1,
			Name:    "JavaScript",
			Price:   100,
			Password:"123",
			Content: []string{"DSA", "Fullstack"},
		},
		{
			Id:      2,
			Name:    "Python",
			Price:   120,
			Password:"456",
			Content: []string{"DSA", "Backend"},
		},
	}

	jsonData, err := json.MarshalIndent(courses,"","\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
}
