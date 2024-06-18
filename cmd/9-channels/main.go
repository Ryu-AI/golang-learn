package main

import (
	"fmt"
	"time"
)

func main() {
	var c = make(chan int) //adding buffer of 5 here makes process exit first then add values
	go process(c)
	// fmt.Println(<-c)
	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}

func process(c chan int) {
	// c <- 123
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Exiting process")
}
