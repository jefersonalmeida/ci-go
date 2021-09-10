package main

import (
	"fmt"
)

func main() {
	fmt.Println(Sum(10, 10))
}

func Sum(a int, b int) int {
	return a + b
}

func SumX(a int, b int) int {
	return a + b + a
}

func Sub(a int, b int) int {
	return a - b
}

func Times(a int, b int) int {
	return a * b
}

func Div(a int, b int) int {
	return a / b
}
