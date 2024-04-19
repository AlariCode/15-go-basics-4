package main

import "fmt"

func main() {
	a := 5
	pointerA := &a
	res := double(a)
	fmt.Println(pointerA)
	fmt.Println(res)
}

func double(num int) int {
	return num * 2
}
