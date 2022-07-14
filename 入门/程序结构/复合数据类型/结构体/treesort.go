package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func add(root *tree, value int) *tree {
	if root == nil {
		t := new(tree)
		t.value = value
		return t
	}
	if value < root.value {
		root.left = add(root.left, value)
	} else {
		root.right = add(root.right, value)
	}
	return root
}

func valuesAppend(root *tree, values []int) []int {
	if root != nil {
		values = valuesAppend(root.left, values)
		values = append(values, root.value)
		values = valuesAppend(root.right, values)
	}
	return values
}

// sort values in place
func sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	valuesAppend(root, values[:0])
}
func main() {
	b := []int{2, 3, 4, 1, 3, 5, 1, 35, 7, 8, 2}

	fmt.Println(b)

	sort(b)
	fmt.Println(b)
}
