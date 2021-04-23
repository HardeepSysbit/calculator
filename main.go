package main

import "fmt"

func main() {

	fmt.Println(add(3, 2))
	fmt.Println(minus(3, 1))

}

func add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}
