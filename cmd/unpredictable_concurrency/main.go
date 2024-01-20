package main

import (
	"fmt"
	"time"
)

func main() {
	go say("hello")
	go say("world")
	time.Sleep(time.Second)
}

func say(toSay string) {
	fmt.Println(toSay)
}
