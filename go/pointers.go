package main
import (
  "fmt"
)
func main() {
  i := 7
  fmt.Println(&i)
}
func inc(x *int) {
  *x++
}
