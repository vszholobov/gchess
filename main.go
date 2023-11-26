package main

import "fmt"

func main() {
	arr := []int{1, 2}
	for i, val := range arr {
		fmt.Println(i, val)
	}
	fmt.Println("Hello world!")
}
