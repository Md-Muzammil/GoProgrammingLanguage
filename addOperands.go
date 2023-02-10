package main

import "fmt"

func main() {

	var num1, num2 int

	fmt.Print("Enter two integers:\n ")
	fmt.Scanf("%d%d%d", &num1, &num2)
	sum := num1 + num2

	fmt.Printf("The sum is : = %d", sum)
}
