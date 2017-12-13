package main

import (
	"fmt"
	"time"
)

type node struct {
	name string
}

func (p *node) print() {
	fmt.Println(p.name)
}

//5 4 3 2 1
func printNumber() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d", i)
	}
}

//5 5 5 5 5, 延迟函数的参数会被求值，但是延迟函数调用表达式不会被求值
func printNumbers() {
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Printf("%d", i)
		}()
	}
}

//参数会求值的情况
func begin(name string) string {

	fmt.Printf("Enter function %s.\n", name)
	return name
}

//end 结束求值
func end(name string) string {
	fmt.Printf("End function %s.\n", name)
	return name
}

//record 记录求值信息
func record() {
	defer end(begin("record"))
	fmt.Println("recording .....")
}

func main() {
loop:
	for {
		switch {
		case true:
			fmt.Println("break ...")
			//跳出循环
			break loop
		}
	}

	//打印是随机的
	data := []*node{{"one"}, {"two"}, {"three"}}
	for _, v := range data {
		go v.print()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("data!")
	//one two three , v属于这个协程的
	data2 := []node{{"one"}, {"two"}, {"three"}}
	for _, v := range data2 {
		v := v
		go v.print()
	}
	time.Sleep(3 * time.Second)

	dd := 1
	dd++
	//这个后面执行
	defer fmt.Println("result =>", func() int { return dd * 2 }()) //print 4，这个在编译的时候已经确定了，使用的是2
	//后面的先执行
	defer func() {
		fmt.Println("dd: ", dd) //print 3,dd是有运行的时候确定的值
	}()
	dd++
	fmt.Println("exit!")
}
