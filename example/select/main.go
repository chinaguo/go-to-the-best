package main

import (
	"fmt"
)

//select 随机选择一个执行，如果两个都到的话
func run(c, q chan int) {
	x := 0
	for {
		select {
		case c <- x:
			x = x + 1
		case <-q:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	q := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {

			fmt.Println("chan is:", <-c)
		}
		q <- 0

	}()

	run(c, q)
}
