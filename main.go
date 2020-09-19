package main

import "fmt"

func main()  {
	var firstName string = "john"

    var lastName string
    lastName = "wick"

    fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Printf("halo %s", firstName, lastName)
	fmt.Println()
	fmt.Printf("halo", firstName, lastName)
	//
	fmt.Printf("halo john wick!\n")
	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName + "!")

	var firstName2 string = "john"
	lastName2 := "wick"

	fmt.Printf("halo2 %s %s!\n", firstName2, lastName2)

}