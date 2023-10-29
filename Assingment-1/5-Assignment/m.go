package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{1, 2, 3, 4, 5}
	if len(slice1) == len(slice2){
		for i := 0; i<len(slice1); i++{
			if slice1[i] != slice2[i]{
				fmt.Println("Slices are Different!")
				return 
			}
		}
		fmt.Println("Slices are Same!")
	}else{
		fmt.Println("Slices are Differnt!")
	}	
	
}
