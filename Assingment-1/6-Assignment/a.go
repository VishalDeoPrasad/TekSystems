package main
import "fmt"

func main() {
	studentGrades := make(map [string]float64)
	studentGrades["Vishal"] = 9.8
	studentGrades["Chinu"] = 9.7
	studentGrades["Rahul"] = 9.6

	for k,v := range studentGrades{
		fmt.Println(k,v)
	}

	delete(studentGrades, "Chinu")
	fmt.Println("Ater Deletion Operation:")
	for k,v := range studentGrades{
		fmt.Println(k,v)
	}
}