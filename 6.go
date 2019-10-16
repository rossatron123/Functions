package main

import "fmt"

func main() {
	h := foo()
	fmt.Println(h)
	x,y := bar()
	fmt.Println(x,y)
}


func foo()int{

	return 42
}

func bar()(int, string){

	return 23, fmt.Sprint("Hello world")
}
