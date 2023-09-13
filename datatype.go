package main

import "fmt"

func main() {
	var a string = "start"
	fmt.Println(a)

	var b int = 5
	fmt.Println(b)

	var d bool
	fmt.Println(d) // 초기화 안한 bool은 false
	d = !d
	fmt.Println(d)

	var e int
	fmt.Println(e) // 초기화 안한 int는 zero

	f := "stop"
	fmt.Println(f)

}
