package main

import (
	"fmt"
)

func def() {
	defer fmt.Println("ok 1")
	defer recover()
	defer fmt.Println("ok2")

	panic("error")

}

// result is 1
func f() (result int) {
	defer func() {
		result++
	}()

	fmt.Printf("%d", 1)

	return 0
}

// result is 5
func f1() (r int) {
	fmt.Println("ff")
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

// result is 1
func f2() (r int) {

	defer func(r int) {
		r = r + 5
	}(r)
	fmt.Println("dd")
	return 1
}

func main() {

	def()

}
