package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func update(s *Student) {
	s.Age += 1
}
func main() {
	s1 := Student{
		Name: "Vishal Kumar",
		Age:  27,
	}
	fmt.Println(s1)
	update(&s1)
	fmt.Println(s1)


}
