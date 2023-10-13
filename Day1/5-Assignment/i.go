package main

import (
	"fmt"
)

func main() {
	numbers := []int{5, 4, 70, 45, 65, 34, 20}
	fmt.Println("Number Before Doubled : ", numbers)

	for i := 0; i < len(numbers); i++ {
		numbers[i] = numbers[i]*2
	}

	fmt.Println("Number After Doubled : ", numbers)


}
