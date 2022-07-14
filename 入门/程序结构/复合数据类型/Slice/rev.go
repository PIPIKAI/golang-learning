package main

import "fmt"

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	reverse(a[1:6])
	reverse(a[:])
	fmt.Println(a)

	b := [...]int{1, 2, 3, 4, 5, 6, 7}
	reverse(b[:])
	fmt.Println(b)

}
