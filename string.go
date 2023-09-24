package main

import "fmt"

func main() {
	// Raw String Literal
	var rawstr string = `안녕 \n반가워` // 안녕 \n반가워 라고 출력(탈출문자도 문자그대로 인식)

	// Interpreted String Literal
	var interstr string = "그럼 \n안녕" // \n을 탈출문자로 인식(한줄 띄움)

	var plusstr string = "기억할게." + "또 보자\n" + "123일 후에"

	fmt.Println(rawstr)
	fmt.Println()
	fmt.Println(interstr)
	fmt.Println()
	fmt.Println(plusstr)
}
