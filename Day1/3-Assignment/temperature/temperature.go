package temperature

func Fahrenheit2Celsius(F float64) float64 {
	c := (F - 32) * 5 / 9
	return c
}
