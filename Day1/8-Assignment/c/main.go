package main

import (
	"go-training/Go-Assignment/TekSystems/Day1/8-Assignment/c/model"
	"go-training/Go-Assignment/TekSystems/Day1/8-Assignment/c/person"
)

func main() {
	p1 := model.Person{
		Name: "Vishal",
		Age : 27,
	}
	person.PrintPerson(p1)
}