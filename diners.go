package main

import (
	"fmt"
	"time"
	"hash/fnv"
	"math/rand"
)

var ph = []string{"Kant", "Thoreau", "Marx"}

type fork byte

const hunger = 3                
const think = time.Second / 100 
const eat = time.Second / 100   

var done = make (chan bool)


func ph_sleep (t time.Duration, rg *rand.Rand) {
	time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
}


func philosopher(name string, 
	dominant_hand, other_hand chan byte,
	done chan bool) {
	fmt.Printf("%s seated\n", name)
	h := fnv.New64a()
	h.Write([]byte(name))
	rg := rand.New(rand.NewSource(int64(h.Sum64())))

	for h := hunger; h!=0; h-- {
		fmt.Println(name, " Hungry")
		<- dominant_hand
		<- other_hand
		fmt.Println(name, " Eating")
		ph_sleep(eat, rg)
		dominant_hand <- 'f'
		other_hand <- 'f'
		fmt.Println(name, " Thinking")
		ph_sleep(think, rg)
	}
	fmt.Println(name, " Done")
	done <- true
	fmt.Println(name, " Left the table")
}


func main() {
	fmt.Println("This is a test")
	place0 := make(chan byte, 1)
	place0 <- 'f'

	placeLeft := place0

	var i int
	for i = 1; i<len(ph);i++ {
		placeRight := make(chan byte, 1)
		placeRight <- 'f'
		go philosopher(ph[i], placeLeft, placeRight, done)
		placeLeft = placeRight
	}
	go philosopher(ph[0], place0, placeLeft, done)

	for _ = range ph {
		<- done
	}

	fmt.Println("Table empty")
}
