package main

import "fmt"

// import "github.com/h31029/alogrithm/array"

func main() {
    defer func() {
        err := recover()
        if err != nil {
            fmt.Println("panic:", err)
        }
    }()
 
    panic("hello, panic!")
}