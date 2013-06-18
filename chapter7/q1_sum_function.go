package main

import "fmt"

func sum(nums []int) int {
    total := 0
    for _, v := range nums {
        total += v
    }
    return total
}
func main() {
    nums := []int{1, 1, 20, 5}
    fmt.Println(sum(nums))
}