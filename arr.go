package main

import "fmt"

func main() {
	// var primes = [6]int{2, 3, 5, 7, 11, 13}
	// var q []int
	// q = primes[2:5]
	// fmt.Println(q)
	// fmt.Println("Length ", len(q))
	// fmt.Println("Capacity ", cap(q))

	t := make([]int, 3, 10)
	fmt.Println("Length ", len(t))
	fmt.Println("Capacity ", cap(t))
	t = t[:5]
	fmt.Println("Length ", len(t))
	fmt.Println("Capacity ", cap(t))
	fmt.Println(t)
	t = append(t, 10)
	fmt.Println("Length ", len(t))
	fmt.Println("Capacity ", cap(t))
	t = append(t, 10)
	fmt.Println("Length ", len(t))
	fmt.Println("Capacity ", cap(t))

	t = append(t, 10)
	fmt.Println("Length ", len(t))
	fmt.Println("Capacity ", cap(t))
	t = append(t, 10)
	fmt.Println("Length ", len(t))
	fmt.Println("Capacity ", cap(t))
	t = append(t, 10)
	fmt.Println("Length ", len(t))
	fmt.Println("Capacity ", cap(t))
	t = append(t, 10)
	fmt.Println("Length ", len(t))
	fmt.Println("Capacity ", cap(t))
	t = append(t, 10)
	fmt.Println("Length ", len(t))
	fmt.Println("Capacity ", cap(t))
	}
