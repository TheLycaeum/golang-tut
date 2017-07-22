package main

import ("fmt"
	"math/rand"
)

func estimate_pi(iters int) float32{
	var inside int
	for i:=1; i<=iters; i++ {
		x := rand.Float32()
		y := rand.Float32()
		if x*x + y*y <= 1 {
			inside++
		}
	}
	return 4 * float32(inside)/float32(iters)
}

func main() {
	fmt.Println(estimate_pi(10000000))
}


