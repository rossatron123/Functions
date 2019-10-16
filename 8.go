package main

import "fmt"

type person struct{

	first string
	last string
	age int
}

func (q person) speak(){

	fmt.Println("Hello there,", q.first, q.last, "I heard you're ", q.age)
}

func main() {

	p := person{

		first: "Ross",
		last: "O'Brien",
		age: 26,
	}
	fmt.Println(p)
	p.speak()
}
