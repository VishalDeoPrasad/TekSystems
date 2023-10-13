package main
import (
	"golang/temperature"
	"fmt"
)
func main() {
	var f float64
	fmt.Print("Enter Temperature in Fahrenheit:")
	fmt.Scanln(&f)
	ans := temperature.Fahrenheit2Celsius(f)
	result := fmt.Sprintf("%.2f",ans)
	fmt.Println(result)
}
