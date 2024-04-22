package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	reverse(&a) // Меняет порядок элементов на обратный
	fmt.Println(a)
}

func reverse(arr *[4]int) {
	for index, value := range *arr {
		(*arr)[len(arr)-1-index] = value
	}
}
