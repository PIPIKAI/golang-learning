package main

import "fmt"

func demo() {
	var s_list []rune
	for _, c := range "Hello,世界" {
		s_list = append(s_list, c)
	}
	fmt.Printf("%q", s_list) // "['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']"
}
func appendInt(a []int, b int) []int {
	var c []int
	clen := len(a) + 1
	if clen <= cap(a) {
		c = a[:clen]
	} else {
		ccap := len(a) * 2
		c = make([]int, clen, ccap)
		copy(c, a)
	}
	c[len(a)] = b
	return c
}
func main() {
	a := []int{1, 2, 3}
	for i := 4; i < 10; i++ {
		a = appendInt(a, i)
		s := fmt.Sprintln(a)
		fmt.Printf("%s  %d\n", s, cap(a))
	}
}
