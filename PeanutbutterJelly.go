package main

import (
	"fmt"
	"time"
)

func main() {
	then := time.Now()

	var ingredient1 string
	var ingredient2 string

	channel := make(chan string)

	go getJelly(channel)
	go getPB(channel)

	ingredient1 = <-channel
	ingredient2 = <-channel

	fmt.Printf("I got some %s and %s to make myself a scrummy sandwhich and it only took me %v time to get it!\n", ingredient1, ingredient2, time.Since(then))
}

func getJelly(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "Jelly"
}

func getPB(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "Peanut Butter"
}
