package main

import "fmt"

func half(num int) (float64, bool) {
    even := false
    if num % 2 == 0 {
        even = true
    } 

    return float64(num) / 2, even
}

func main() {
    fmt.Println(half(23.0))
}