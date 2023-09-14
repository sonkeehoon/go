package main

import "fmt"

func main() {

	const a int = 1

	const b, c int = 5, 10
	const d, e = 15, 20

	const hello = "Hello"

	fmt.Println(a)
	fmt.Println(b, c)
	fmt.Println(d, e)
	fmt.Println(hello)

	const (
		val1 = 1
		val2 = "hello"
	)

	const (
		c1 = 10 // 첫번째 값은 초기화 필수
		c2
		c3
		c4 = "Hello" // 다른 타입도 가능
		c5
		c6 = iota // c9까지 index값 저장
		c7
		c8
		c9
		c10 = "Bye"
		c11
	)
	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11)

}
