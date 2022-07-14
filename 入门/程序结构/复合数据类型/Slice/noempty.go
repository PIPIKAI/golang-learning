package main

import "fmt"

func noempty(s []string) []string {
	var new_s []string
	for _, c := range s {
		if c != "" {
			new_s = append(new_s, c)
		}
	}
	return new_s
}
func main() {
	a := []string{"hello ", "", "world", "", "!"}

	fmt.Printf("%q \n %q", a, noempty(a))
}
