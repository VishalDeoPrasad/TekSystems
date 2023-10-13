package main

import "fmt"

type Animal interface {
	MakeSound() string
}

type Dog struct { Name string }
type Cat struct { Name string }

func (d Dog) MakeSound() string {
	fmt.Println("Who is making noise:", d.Name)
	return ""
}

func (c Cat) MakeSound() string {
	fmt.Println("Who is making noise:", c.Name)
	return ""
}
func main() {
	a1 := Dog{
		Name:"Tom",
	}
	a2 := Cat{
		Name:"Bubble",
	}
	a1.MakeSound()
	a2.MakeSound()
}