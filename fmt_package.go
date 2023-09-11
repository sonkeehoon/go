package main

import "fmt"

func main() {
	var num1 int = 1
	var num2 int = 2

	fmt.Print("This is Number: ")
	fmt.Print(num1)
	fmt.Print(num2)

	fmt.Print("\n\n")

	fmt.Println("This is Number: ")
	fmt.Println(num1)
	fmt.Println(num2)

	fmt.Printf("This is Number: ")
	fmt.Printf("%d%d", num1, num2)
}
