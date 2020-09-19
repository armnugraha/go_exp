package library

import "fmt"

// huruf pertama besar = public
// func SayHello() {
// add param name for call introduce
func SayHello(name string) {
	fmt.Println("hello")
	introduce(name)
}

// huruf pertama kecil = private
func introduce(name string) {
    fmt.Println("nama saya", name)
}