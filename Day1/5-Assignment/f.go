package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	data := os.Args[1:]
	k, _ := strconv.Atoi(data[0])
	numbers := []int{5, 20, 15, 1, 2}
	for i := 1; i < len(numbers); i++ {
		if numbers[i] == k {
			fmt.Println("Number Found:")
			return 
		}
	}
	fmt.Println("Number not found")
}
