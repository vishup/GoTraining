package main

import (
	"fmt"
	"time"
)

//pinger can only send/write to a channel
func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

//printer can only receive/read from a channel
func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	c := make(chan string)

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
