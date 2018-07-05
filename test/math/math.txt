package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func mult(w int, z int) int{
	return w * z
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(mult(12, 12))
}