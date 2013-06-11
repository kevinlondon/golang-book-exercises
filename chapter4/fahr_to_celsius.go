package main

import "fmt"

func main() {
    fmt.Print("Enter a temperature (in Fahrenheit): ")
    var temp float64
    fmt.Scanf("%f", &temp)

    celsius := (temp - 32) * 5 / 9

    fmt.Println("Celsius:", celsius)
}