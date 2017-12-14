package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(l []int, left, right int) {
	if left < right {
		q := partition(l, left, right)
		if q > left {
			quickSort(l, left, q-1)
		}

		if q < right {
			quickSort(l, q+1, right)

		}
	}
}

//分治的策略，可以进行随机取x
func partition(l []int, left, right int) int {
	x := l[right]
	i := left - 1

	for j := left; j < right; j++ {
		if l[j] <= x {
			i = i + 1
			l[i], l[j] = l[j], l[i]
		}
		fmt.Printf("i:%d, j:%d, x:%d, l:%v. \n", i, j, x, l)
	}

	l[i+1], l[right] = l[right], l[i+1]
	fmt.Println("l", l)
	return i + 1
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
	time.Sleep(1 * time.Second)

	var z []int
	for i := 0; i < 10; i++ {

		z = append(z, rand.Intn(1000))
	}
	fmt.Println(z)
	quickSort(z, 0, len(z)-1)
	fmt.Println(z)
}
