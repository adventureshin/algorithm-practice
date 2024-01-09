package main

import (
    "bufio"
    "fmt"
    "os"
    "errors"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    scanner.Scan()
    num, _ := strconv.Atoi(scanner.Text())
    var stack []int

    for i := 0; i < num; i++ {
        scanner.Scan()
        commandType := scanner.Text()
        switch commandType {
        case "push":
            scanner.Scan()
            commandValue, _ := strconv.Atoi(scanner.Text())
            stack = append(stack, commandValue)
        case "pop":
            if len(stack) == 0 {
                fmt.Println(-1)
            } else {
                index := len(stack) - 1
                fmt.Println(stack[index])
                stack = stack[:index]
            }
        case "size":
            fmt.Println(len(stack))
        case "empty":
            if len(stack) == 0 {
                fmt.Println(1)
            } else {
                fmt.Println(0)
            }
        case "top":
            if len(stack) == 0 {
                fmt.Println(-1)
            } else {
                fmt.Println(stack[len(stack)-1])
            }
        default:
            panic(errors.New("Invalid command"))
        }
    }
}

