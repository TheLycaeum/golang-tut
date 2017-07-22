package main

import ("fmt"
	"math/rand")


func estimate_pi(iters int) float32 {
	inside := 0
	for i:=1; i<iters+1; i++ {
		x := rand.Float32()
		y := rand.Float32()
		if x*x + y*y <=1 {
			inside += 1
		}
	}
	return 4.0*float32(inside)/float32(iters)
}

func main() {
	pi := estimate_pi(10000000)
	fmt.Println(pi)
}	
