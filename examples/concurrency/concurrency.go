package main

import (
	"fmt"
	"time"
)

func reportForDuty(name string, laziness int, ch chan string) {
	time.Sleep(time.Second * time.Duration(laziness))
	ch <- name
}

func main() {
	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("...no")
	}()
	fmt.Println("Are you awake?")
	time.Sleep(time.Second * 2)

	ch := make(chan bool)
	go func(ch chan bool) {
		time.Sleep(time.Second * 2)
		fmt.Println("...no")
		ch <- true
	}(ch)
	fmt.Println("Are you awake?")
	<-ch
	fmt.Println("Stop being lazy.")

	first := make(chan string)
	go reportForDuty("Jimmy", 3, first)
	second := make(chan string)
	go reportForDuty("Timmy", 1, second)

	var output string
	for i := 0; i < 2; i++ {
		select {
		case name := <-first:
			output += name
		case name := <-second:
			output += name
		}
	}
	fmt.Println(output)
}
