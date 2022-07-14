package main

import "fmt"

type point struct {
	x, y int
}
type circle struct {
	point
	radius int
}
type wheel struct {
	circle
	spokes int
}

func main() {
	w := wheel{circle: circle{point: point{1, 2}, radius: 5}, spokes: 10}
	fmt.Printf("%#v", w)
}
