package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("main goroutine")
    for i := 0; i < 10; i++ {
        fmt.Printf("i")
        time.Sleep(1 * time.Second)
    }
    go func() {
        time.Sleep(2 * time.Second)
        panic("fuck")
    }()
    time.Sleep(3 * time.Second)
}

