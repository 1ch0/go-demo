package main

import "fmt"

func main() {
	t := &T{}
	t.Print()
	t.Print2()
}

type T struct{}

func (t T) Print() {
	fmt.Println("T: ", t)
}

func (t T) Print2() {
	fmt.Println("T2: ", t)
}
