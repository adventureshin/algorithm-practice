package main

import (
  "fmt"
)

func main() {
  var a int
  var b int
  var c []int
  var d int
  fmt.Scan(&a)
  for i := 0; i < a; i++ {
    fmt.Scan(&b)
    c = append(c, b)
  }
  fmt.Scan(&d)
  var sum int
  for i := 0; i < a; i++ {
    if c[i] == d {
      sum++
    }
  }
  fmt.Println(sum)
}
