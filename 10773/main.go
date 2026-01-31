package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanWords)
  scanner.Scan()
  n, _ := strconv.Atoi(scanner.Text())
  var sum_n int
  var stack []int
  for i := 0; i < n; i++ {
    scanner.Scan()
    a, _ := strconv.Atoi(scanner.Text())
    switch a {
    case 0:
      sum_n -= stack[len(stack)-1]
      stack = stack[:len(stack)-1]
    default:
      sum_n += a
      stack = append(stack, a)
    }
  }
  fmt.Println(sum_n)
}
