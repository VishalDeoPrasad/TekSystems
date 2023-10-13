package main

import "fmt"

type Ractangle struct {
	Width  int
	Height int
}

func area(dim Ractangle) (int, int) {
	return dim.Width * dim.Height, 2 * (dim.Height + dim.Width)
}
func main() {
	d := Ractangle{
		Width:  10,
		Height: 20,
	}
	area, paremeter := area(d)
	fmt.Println("Area:",area)
	fmt.Println("Paremeter:",paremeter)
}