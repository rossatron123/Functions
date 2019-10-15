package main


import "fmt"

func main() {

	foo()
	s := bar("Jason")
	fmt.Println(s)
	a, b, c := woo("Ross", "O'Brien", 26)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}

func foo(){
	fmt.Println("Hi Rachel!")
}

func bar(s string)string{
	a := fmt.Sprintln("Hi there, ", s)
	return a
}

func woo(s string, t string, a int)(string, int, bool){

	q := fmt.Sprintln("Hey there ",s,t )
	w := 26
	e := true

	return q,w,e
}


