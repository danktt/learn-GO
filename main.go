package main

import "fmt"

func main() {

	x := 10
	y := "User name"


	fmt.Printf("%v %T \n", x, x)
	fmt.Printf("%v %T \n", y, y)

	_, error := fmt.Println("Hello, GO!! ")
		fmt.Println( error)
}
