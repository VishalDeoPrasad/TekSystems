package calculator
import (
	"fmt"
)

func Square(n int)int  {
	return n*n
}

func Sum(n1, n2 int)int{
	return n1+n2
}

func Diff(n1, n2 int)int{
	return n1-n2
}

func Product()int{
	var n1, n2 int
	fmt.Print("Enter First Number:")
	fmt.Scanln(&n1)
	fmt.Print("Enter Second Number:")
	fmt.Scanln(&n2)
	fmt.Printf("%dx%d = ",n1,n2)
	return n1*n2
	
}

func QuotientRemainder(n1, n2 int)(int, int){
	rem := n1%n2
	quotient := n1/n2
	return quotient, rem
}
