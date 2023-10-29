package person

import (
	"fmt"
	"go-training/Go-Assignment/TekSystems/Day1/8-Assignment/c/model"
)

func PrintPerson(p model.Person) {
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}
