package main

import ("fmt"
	"time")

func say(s string) {
	for i:=0;i<5;i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
		}
	}

func squares(ip chan int, op chan int) {
	var x int
	fmt.Println("I am starting")
	for { 
		x = <- ip
		fmt.Println("Got  ", x)
		op <- x*x
		fmt.Println("Sent ", x*x)

	}
}

	

func main() {
	send := make(chan int)
	recieve := make(chan int)
	go squares(send, recieve)
	for i:=0; i<10; i++ {
		send <- i
		fmt.Println(i, <- recieve)
	}
}




