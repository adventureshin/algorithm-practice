package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func main() {
  scanner.Split(bufio.ScanWords)
  scanner.Scan()
  num, _ := strconv.Atoi(scanner.Text())
  numMap := make(map[string]bool)
  for i := 0; i < num; i++ {
    scanner.Scan()
    numMap[scanner.Text()] = true
  }
  scanner.Scan()
  num, _ = strconv.Atoi(scanner.Text())
  for i := 0; i < num; i++ {
    scanner.Scan()
    if _, ok := numMap[scanner.Text()]; ok {
      fmt.Println(1)
    } else {
      fmt.Println(0)
    }
  }
}
