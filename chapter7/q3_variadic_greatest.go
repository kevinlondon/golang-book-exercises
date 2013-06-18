package main

import "fmt"

func greatest(args ...int) int {
    highest_num := 0
    for _, v := range args {
        if v > highest_num {
            highest_num = v
        }
    }
    return highest_num
}

func main() {
    fmt.Println("Greatest: ", greatest(20, 40, 8, 16, 3))
}