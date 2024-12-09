package main

import "fmt"

// import "github.com/h31029/alogrithm/array"

func helloPanic()  {
    defer func() {
        err := recover()
        if err != nil {
            fmt.Println("panic:", err)
        }
    }()
 
    panic("hello, panic!")
}

func main() {
    arr := [10]int{0,1,2,3,4,5,6,7,8,9}
    s1 := arr[1:4]
    println("len:", len(s1), "cap:", cap(s1))
    s1 = append(s1, 10)
    println("after expansion: len:", len(s1), "cap:", cap(s1))
    for _, v := range arr {
        println(v)
    }
}