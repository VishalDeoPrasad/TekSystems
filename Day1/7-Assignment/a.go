package main

import "fmt"

type Employee struct {
	Name string
	Age  int
	City string
}

func main() {
	e1 := Employee{
		Name: "Vishal",
		Age:  26,
		City: "Banglore",
	}

	e2 := Employee{
		Name: "Kumar",
		Age:  27,
		City: "Pune",
	}

	fmt.Println(e1)
	fmt.Println(e2)

}