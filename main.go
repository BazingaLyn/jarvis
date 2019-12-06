package main

import (
	"fmt"
	"time"
)

func main() {
	var Ball int
	table := make(chan int)
	go player("1", table)
	go player("2", table)

	table <- Ball
	time.Sleep(1 * time.Second)
	<-table
}

func player(name string, table chan int) {
	for {
		fmt.Println("name is :" + name)
		ball := <-table
		ball++
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
