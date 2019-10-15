package main

import "fmt"

type person struct{
	first string
	last string
}

type sagent struct{

	person
	ltk bool
}

func (a sagent) foo(){

	fmt.Println("HXXXXX there:",a.first, a.last)
}

func main() {

	p1 := sagent{
		person: person{
			first: "Ross",
			last: "O'Brien",
		},
		ltk: true,
	}

	s1 := sagent{
		person: person{
			first: "Sara",
			last: "Jones",
		},
		ltk: false,
	}

	fmt.Println(p1)
	fmt.Println(s1)

	s1.foo()
	p1.foo()


}
