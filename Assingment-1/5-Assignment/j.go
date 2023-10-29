package main

import (
	"fmt"
)

func main() {
	slice1 := []int{65, 34, 20}
	slice2 := []int{5, 4, 70}
	slice2 = append(slice2, slice1...)
	fmt.Println(slice2)
	
	
}
