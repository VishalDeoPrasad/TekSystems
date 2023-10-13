package main

import (
	"fmt"
	"golang/calculator"
)

func main() {
	fmt.Println("Square of 10:", calculator.Square(10))
	fmt.Println("10+20 = ", calculator.Sum(10, 20))
	fmt.Println("100-20 = ", calculator.Diff(100, 20))
	fmt.Println(calculator.Product())
	q, r := calculator.QuotientRemainder(65, 15)
	fmt.Printf("Quotient=%d, Remainder=%d", q, r)
}
