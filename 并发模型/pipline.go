// Created by Hisen at 2019-07-11.
package main

import (
	"fmt"
)

func c() {
	for i := 0; i < 100000; i++ {
		ch1 <- i
	}
	close(ch1)
}

func m() {
	for i := range ch1 {
		ch2 <- i * 2
	}
	close(ch2)
}

var ch1 = make(chan int, 100)
var ch2 = make(chan int, 100)

func main() {
	go c()
	go m()
	for i := range ch2 {
		fmt.Println(i)
	}
}
