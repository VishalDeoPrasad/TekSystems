package main

import (
	"fmt"
)

func matchValue(slice []int, key int)bool{
	for i := 0; i<len(slice); i++{
		if slice[i] == key{
			return false
		}
	}
	return true
}

func main() {
	slice := []int{1,2,3,1,2,3,5,6,6,2}
	list := []int{}
	list = append(list, slice[0])
	for i := 1; i<len(slice); i++{
		if matchValue(list, slice[i]){
			list = append(list, slice[i])
		}
	}
	fmt.Println(list)
}
