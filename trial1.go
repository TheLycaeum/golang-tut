package main


import "fmt"

type Point struct{
	X int
	Y int
	}

func main() {
	p0 := Point{1, 2}
	fmt.Println(p0.X, p0.Y)
	
	var p_i *Point
	p_i = &p0
	fmt.Println((*p_i).X)
	fmt.Println(p_i.X)

	
	}
