package main

import(
 "fmt"
 "sync"
)

func calc(a, b int) int{
  res := a+b
  fmt.Println(a, b, res)
  return res
}

func main() {

  a := 0
  b := 1

  defer calc(a, calc(a, b))

  a = 3
  b = 4

  defer fmt.Println(a, b,  calc(a, b))
  a = 4
  p := new(person)
  p.pe = make(map[string]int)
  // p.mu = new(sync.Mutex)
  for i := 0; i < 5; i++ {
     p.Add("1", 1)
     fmt.Println("person", p.pe["1"])
  }

}

type person struct {
   mu sync.Mutex
   pe map[string]int
}

func (p *person) Add(name string, age int) {
   p.mu.Lock()
   p.pe[name] = age
   p.mu.Unlock()
}

func (p *person) Get(name string) int {

   return p.pe[name]

}