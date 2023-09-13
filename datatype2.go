package main

import "fmt"

func main() {
	// 여러개의 변수 한번에 선언
	var a, b int = 10, 20
	fmt.Println(a, b)

	x, y, z := 1, 2, 3
	fmt.Println(x, y, z)

	var str1, str2 string = "Hello", "World"

	fmt.Println(str1, str2)

}
