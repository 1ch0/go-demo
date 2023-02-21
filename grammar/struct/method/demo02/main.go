package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	var i MyInt = 1
	fmt.Println(i.string())
	fmt.Println(unsafe.Pointer(&i))
	fmt.Println(i.multiply())
	fmt.Println(unsafe.Pointer(&i))
}

type MyInt int

func (m MyInt) string() string {
	return strconv.Itoa(int(m))
}

func (m MyInt) multiply() MyInt {
	return m * m
}
