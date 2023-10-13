package main
import (
	"fmt"
	"math/rand"
)
var random_no, guess int

func main() {
	random_no = rand.Intn(10)
	for {
		fmt.Print("Try Guessing:")
		fmt.Scanln(&guess)
		if guess == random_no{
			fmt.Println("You guessed it right:")
			return 
		}else{
			if guess > random_no{
				fmt.Print("too high, ")
			}else{
				fmt.Print("too loo, ")
			}	
			
		}
	}
}