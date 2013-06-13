package main

import "fmt"

func main() {
    elements := map[string]string {
        "H": "Hydrogen",
        "He": "Helium",
        "Li": "Lithium",
        "Be": "Beryllium",
        "B": "Boron",
        "C": "Carbon",
        "N": "Nitrogen",
        "O": "Oxygen",
        "F": "Fluorine",
        "Ne": "Neon",
    }

    if el, ok := elements["Li"]; ok {
        fmt.Println(el)
    }
}