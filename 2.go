package main

import "fmt"

func main() {

	foo(1,2,3,4,5,6,7)

}


func foo(x ...int){

	fmt.Println(x)

	sum := 0

	for i, v := range x{
		sum += v
		fmt.Println("The index position,",i, " with the value:",v, "and the sum is: ",sum)
	}
}
