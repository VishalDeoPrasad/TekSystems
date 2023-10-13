package main

import "fmt"

func main() {
	studentData := map[string]map[any]any{
		"Vishal": {
			"Age":   26,
			"Grade": 9.9,
			"City":  "Banglore",
		},
		"Rahul": {
			"Age":   25,
			"Grade": 9.7,
			"City":  "Delhi",
		},
		"Shashi": {
			"Age":   24,
			"Grade": 9.5,
			"City":  "Bihar",
		},
	}
	studentData["Ravi"] = map[any]any{
		"Age":   22,
		"Grade": 9.4,
		"City":  "Pune",
	}
	for k, v := range studentData {
		fmt.Println(k)
		for ik, iv := range v {
			fmt.Println("--->", ik, iv)
		}
	}
}
