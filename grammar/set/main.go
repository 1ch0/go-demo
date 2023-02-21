package main

import "fmt"

func main() {
	s := make(set)
	s.add(1)
	s.add(2)
	s.add(3)
	fmt.Println(s)
	s.remove(2)
	fmt.Println(s)
	fmt.Println(s.exist(1))
}

type set map[int]struct{}

func (s set) add(num int) {
	s[num] = struct{}{}
}

func (s set) remove(num int) {
	delete(s, num)
}

func (s set) exist(num int) bool {
	_, ok := s[num]
	return ok
}
