package shape

import (
	"go-training/Go-Assignment/TekSystems/Day1/8-Assignment/d/model"
	"math"
)

func AreaOfCircle(r model.Circle) float64 {
	return math.Pi * (r.Radius * r.Radius)
}
