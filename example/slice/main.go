package main

import (
	"fmt"
)

//result:
//[100 101 102 0 0 0 0 0 0 0]
//[120 101 102 0 0 0 0 0 0 0]
//map[x:ok]

//地址值
func change(x []int) {
	x[0] = 120
}

//地址值
func changeMap(m map[string]string) {
	m["x"] = "ok"

}

func main() {

	arr := make([]int, 10)

	arr[0] = 100
	arr[1] = 101
	arr[2] = 102

	fmt.Println(arr)
	change(arr)
	fmt.Println(arr)

	m := make(map[string]string)
	m["x"] = "no"
	changeMap(m)
	fmt.Println(m)
}
