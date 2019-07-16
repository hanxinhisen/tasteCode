// Created by Hisen at 2019-06-19.
package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("sssssssss")
	st := make(chan os.Signal)
	signal.Notify(st, os.Interrupt)
	<-st
	fmt.Println("退出")
}
