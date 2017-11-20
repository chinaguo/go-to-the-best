package main

import (
	"fmt"
	"runtime"
)

func hello(s string) {
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	done := false
	go func() {
		done = true
	}()

	//主动让车时间片
	for !done {
		fmt.Println("Ok")
		runtime.Gosched()
	}

	fmt.Println("done!")
}
