package main

import (
	"fmt"
	"time"
)

func quickSort() {
	i := 0
	i++
}

func dataSlice() {
	data := []int{1, 2, 3}
	i := 0
	i++                  // ++i error
	fmt.Println(data[i]) //data[i++] error
}

func dataOperator() {
	var d uint8 = 2

	fmt.Printf("%08b\n", ^d)
	//fmt.Println(^2) //~2 error
}

func main() {
	//这个是阻塞的channel
	ch := make(chan int)
	done := make(chan struct{})

	for i := 0; i < 3; i++ {
		go func(idx int) {
			select {
			case ch <- (idx + 1) * 2:
				fmt.Println(idx, "sent result")
			case <-done:
				fmt.Println(idx, "exiting")
			}
		}(i)
	}

	//get first result
	fmt.Println("result:", <-ch)
	close(done)

	x := "text"
	xByte := []byte(x)
	xByte[0] = 'T'
	fmt.Println(string(xByte))

	//do other work
	time.Sleep(3 * time.Second)
}
