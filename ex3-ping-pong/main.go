package main

import (
	"fmt"
	"time"
)

func ping(ch chan string) {
	for {
		ch <- "Ping"
	}
}

func pong(ch chan string) {
	for {
		ch <- "Pong"
	}
}

func writeChannel(ch chan string) {
	for {
		message := <-ch
		fmt.Println(message)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	ch := make(chan string)
	go ping(ch)
	go writeChannel(ch)
	go pong(ch)

	var scan string
	fmt.Scanln(&scan)
}
