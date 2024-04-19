package main

import "fmt"

func main() {
	a := 5
	double(&a)
	fmt.Println(a)
}

func double(num *int) {
	*num = *num * 2
}
