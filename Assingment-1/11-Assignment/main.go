package main

import "fmt"

type Address struct {
	Steet   string
	City    string
	ZipCode int
}
type Person struct {
	Name string
	Address
}

func main() {
	p1 := Person{
		Name: "Vishal Kumar",
		Address: Address{
			Steet:   "BYM Layout",
			City:    "Banglore",
			ZipCode: 560017,
		},
	}
	fmt.Printf("Name:%q\nAddress:%q, %q, %d",p1.Name, p1.Steet, p1.City, p1.ZipCode)
}
