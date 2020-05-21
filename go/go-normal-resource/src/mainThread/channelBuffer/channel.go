package main

import (
	"fmt"
	"time"
)

func main() {
	m1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 10)
		m1 <- "Done!"
	}()

	select {
	case passingMessage := <-m1:
		fmt.Println(passingMessage)
	case <-time.After(2 * time.Second):
		fmt.Println("m1 goroutine was.......Time Out......")
	}
}
