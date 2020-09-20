package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}

	return true, nil
}

func main() {
	// A.36.1
	// var input string

	// fmt.Println("Type some number : ")
	// fmt.Scanln(&input)

	// var number int
	// var err error
	// number, err = strconv.Atoi(input)

	// if err == nil {
	// 	fmt.Println(number, " Is number")
	// } else {
	// 	fmt.Println(input, " Is not number")
	// 	fmt.Println(err.Error())
	// }

	// A.36.2
	var name string
	fmt.Print("Type Your Name :")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("Hallo", name)
	} else {
		fmt.Println(err.Error())
	}

}
