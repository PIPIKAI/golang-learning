package main

import (
	"fmt"
	"net/http"
)

func Sort(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
func Mul(a, b int) int {
	return a * b
}
func Conn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
func main() {
	fmt.Println("main()")
}
