package library

import "fmt"

// huruf pertama besar = public
func SayHello() {
    fmt.Println("hello")
}

// huruf pertama kecil = private
func introduce(name string) {
    fmt.Println("nama saya", name)
}