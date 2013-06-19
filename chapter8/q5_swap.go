package main

import "fmt"

func swap(xPtr *int, yPtr *int) {
    *xPtr, *yPtr = *yPtr, *xPtr
}

func main() {
    x := 1
    y := 2
    fmt.Println(x, y) // 1, 2
    swap(&x, &y)
    fmt.Println(x, y) // Becomes 2, 1
}