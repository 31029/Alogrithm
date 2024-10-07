package main

import "fmt"

// import "github.com/h31029/alogrithm/array"

func main()  {
	sn := []int{1, 2}
	sn1 := sn[2:]

	for k, v := range sn1 {
		fmt.Printf("k:%v, v:%v \n", k, v)
	}
}
