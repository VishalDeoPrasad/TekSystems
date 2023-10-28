package main

import "fmt"

func divide(numerator, denominator int) int {
	if denominator == 0 {
		panic("Denominator can not be zero")
	}
	return numerator / denominator
}

func safeDivide(numerator, denominator int) {
	defer func() {
		rec := recover()

		if rec != nil {
			fmt.Println("panic recovered", rec)
			divide(numerator, denominator+1) //testing
		}
	}()

	fmt.Println(divide(numerator, denominator))

}
func main() {
	safeDivide(3, 0)
}
