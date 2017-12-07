package main

import (
	"fmt"
	"math"
)

//这个写法在某种情况是有问题的。
func max(a int64, b int64) int64 {
	return int64(math.Max(float64(a), float64(b)))

}

func main() {
	fmt.Println(math.MaxInt64, math.MaxInt64-1, max(math.MaxInt64, math.MaxInt64-1))
}
