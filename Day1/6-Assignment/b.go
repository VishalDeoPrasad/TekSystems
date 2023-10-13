package main
import "fmt"

func main() {
	studentGrades := make(map [string]int)
	studentGrades["Kiwi"] = 9
	studentGrades["Mango"] = 9
	studentGrades["Orange"] = 10

	for k,v := range studentGrades{
		fmt.Println(k,v)
	}

}