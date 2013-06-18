package main

import "fmt"

func makeOddGenerator() func() uint {
    i := uint(1)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}

func main() {
    nextOdd := makeOddGenerator()
    fmt.Println(nextOdd())      // 1
    fmt.Println(nextOdd())      // 3
    fmt.Println(nextOdd())      // 5
}