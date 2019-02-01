package main

import "fmt"

func main() {
	fmt.Println(point{1,1})
	fmt.Println(equation{1,2,3})
}

type point struct {
	x float64
	y float64
}

type equation struct {
	x, y, c float64
}