package main

import (
	"fmt"
	"math/rand"
)

func mergeSort(r []int) []int {
	length := len(r)
	if length <= 1 {
		return r
	}

	num := length / 2
	left := mergeSort(r[:num])
	right := mergeSort(r[num:])

	return merge(left, right)
}

//这个算法主要使用空间换取时间，有点取巧,此方法来自于网络，我更喜欢c或者c++方式的实现
func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {

		if left[l] < right[r] {
			result = append(result, left[l])
			L++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return
}

func main() {
	var z []int
	for i := 0; i < 10; i++ {
		z = append(z, rand.Intn(1000))
	}

	arr := mergeSort(z)
	fmt.Println(arr)
}
