package main

import(
  "fmt"
)

func def() {
   defer fmt.Println("ok 1")
   defer recover()
   defer fmt.Println("ok2")

   panic("error")

}

func main() {

   def()

}
