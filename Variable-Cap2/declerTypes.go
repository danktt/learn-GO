package main

import "fmt"


type hotdog int
type busscar float64

var b hotdog = 42

func main() {

	x := 10
	fmt.Printf("%v %T \n", x, x)  // 10 int
	fmt.Printf("%v %T \n", b, b)  // 42 main.hotdog

	x = int(b)
	fmt.Printf("%v %T \n", x, x)  // 42 int
}
