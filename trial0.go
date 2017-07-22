package main

import "fmt"


func add(x, y int) int {
	return x + y
}

func greet() {
	defer fmt.Println("World")
	fmt.Println("Hello")
	}

func sumdiff (x, y int) (sum int, diff int) {
	sum = x + y
	diff = x - y
	return
	}

func main() {
	// a := 10
	// b := 5
	// sum := add(a, b)
	// var t float32
	// t = float32(sum)
	// fmt.Println(t)
	// fmt.Println(Pi)
	// for b<10 {
	// 	fmt.Println(b)
	// 	b++
	// }

	// if q:= a*10; q > 100 {
	// 	fmt.Println("Hello")
	// } else {
	// 	fmt.Println("Goodbye")
	// }
	greet()

	var p_i *int
	q := 15
	p_i = &q
	fmt.Println("The value is ",*p_i)
	
}


