package main

import (
    "fmt"
    "strconv"
)

func main() {
    var num1, num2 float64
    var operator string

    fmt.Println("Kalkulator Sederhana")
    fmt.Println("====================")
    fmt.Print("Masukkan angka pertama: ")
    fmt.Scanln(&num1)

    fmt.Print("Masukkan operator (+, -, *, /): ")
    fmt.Scanln(&operator)

    fmt.Print("Masukkan angka kedua: ")
    fmt.Scanln(&num2)

    result := calculate(num1, num2, operator)
    fmt.Printf("Hasil: %v\n", result)
}

func calculate(num1, num2 float64, operator string) float64 {
    var result float64
    switch operator {
    case "+":
        result = num1 + num2
    case "-":
        result = num1 - num2
    case "*":
        result = num1 * num2
    case "/":
        if num2 != 0 {
            result = num1 / num2
        } else {
            fmt.Println("Error: Pembagian dengan nol tidak diizinkan")
            return 0
        }
    default:
        fmt.Println("Error: Operator tidak valid")
        return 0
    }
    return result
}

