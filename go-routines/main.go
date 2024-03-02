package main

import (
	"fmt"
	"time"
)

func say(message string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(message)
	}
}

func main() {
	go say("world")
	say("hello")
}
