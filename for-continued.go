package main

import "fmt"

func main() {
	sum := 7
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
