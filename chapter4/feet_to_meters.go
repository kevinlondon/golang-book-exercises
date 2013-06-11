package main

import "fmt"

func main() {
    fmt.Print("Enter the number of feet: ")
    var feet float64
    fmt.Scanf("%f", &feet)

    // Feet to Meters: 1 f = 0.3048 m
    meters := feet * .3048

    fmt.Println("Meters:", meters)
}