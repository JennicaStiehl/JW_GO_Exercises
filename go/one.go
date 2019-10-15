package main

import (
  "fmt"
  "math"
  "errors"
)
func main() {
  fmt.Println("Hello World!")
  var x int = 5 //or x := 5
  var y int = 7 // y := 7
  var getSum int = x + y // sum := x+y
  result := sum(2, 3)

  fmt.Println(getSum)
  fmt.Println(result)

result2, err := sqrt(16)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(result2)
  }
}

func sum(x int, y int) int {
  return x + y
}

func sqrt(x float64) (float64, error) {
if x < 0 {
  return 0, errors.New("undefined for negative numbers")
}
return math.Sqrt(x), nil
}
//run one.go
//go install
