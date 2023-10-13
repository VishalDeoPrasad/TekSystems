package main

import "fmt"

func main() {
	slice := []string{"Apple", "Banana", "Kiwi", "Orange"}
	for f := range slice {
		fmt.Println(slice[f])
	}
}
