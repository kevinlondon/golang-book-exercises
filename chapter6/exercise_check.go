package main

import "fmt"

func main() {
    smallest := 10000000
    x := []int{
        48,96,86,68,
        57,82,63,70,
        37,34,83,27,
        19,97, 9,17,
    }

    for i := 0; i < len(x); i++ {
        if x[i] < smallest {
            smallest = x[i]
        }
    }

}