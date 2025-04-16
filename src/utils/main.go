package main

import "fmt"

func main() {
	var scores []int = make([]int, 5, 10)
	fmt.Println(scores[:])   // [0 0 0 0 0]
	fmt.Println(scores[1:3]) // 5
}
