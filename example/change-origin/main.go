package main

import "fmt"

type data struct {
	num   int
	key   *string
	items map[string]bool
}

//1
func (d *data) pMethod() {
	d.num = 7
}

//2 和1是有区别的
func (d data) vMethod() {
	d.num = 8
	*d.key = "v.key"
	d.items["vMethod"] = true
}

func main() {
	key := "key.1"
	d1 := data{1, &key, make(map[string]bool)}
	fmt.Printf("num=%v key=%v items=%v\n", d1.num, *d1.key, d1.items)

	d1.pMethod()
	fmt.Printf("num=%v key=%v items=%v\n", d1.num, *d1.key, d1.items)

	d1.vMethod()
	fmt.Printf("num=%v key=%v items=%v\n", d1.num, *d1.key, d1.items)
}
