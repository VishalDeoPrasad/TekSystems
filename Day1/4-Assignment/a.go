package main
import "fmt"
func checkEvenOdd(a int){
	if a%2 == 0{
		fmt.Println("Number is Even:")
	}else{
		fmt.Println("Number is Odd")
	}
}
func main() {
	checkEvenOdd(10)
}