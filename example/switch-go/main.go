package main

import "fmt"

func main() {
loop:
	for {
		switch {
		case true:
			fmt.Println("break ...")
			break loop
		}
	}

	fmt.Println("exit!")
}
