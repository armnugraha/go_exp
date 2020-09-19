package main

import (
    "fmt"
    "reflect"
)

func main() {
    number := 23
    var reflectValue = reflect.ValueOf(number)

    fmt.Println("Tipe Variable adalah : ", reflectValue.Type())

    if reflectValue.Kind() == reflect.Int {
        fmt.Println("Nilai Variable : ", reflectValue.Int())
    }
}