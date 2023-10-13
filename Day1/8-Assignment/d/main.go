package main

import (
	"fmt"
	"go-training/Go-Assignment/TekSystems/Day1/8-Assignment/d/model"
	"go-training/Go-Assignment/TekSystems/Day1/8-Assignment/d/shape"
)

func main() {
	r := model.Circle{
		Radius: 3,
	}
	fmt.Println("Area of Circle:", shape.AreaOfCircle(r))
}
