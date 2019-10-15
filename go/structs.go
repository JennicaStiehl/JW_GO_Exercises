package main

import (
  "fmt"
)
type person struct {
  name string
  age int
}

func main() {
  p := person{name: "joe", age:23}
  fmt.Println(p.age)
}
