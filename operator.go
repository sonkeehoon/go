package main

import "fmt"

func main() {
	var num1, num2 int = 5, 10
	var str1, str2 string = "Hello", "World!"

	// 4칙연산(+, -, *, / 와 %(나머지를 구해주는) 연산)
	fmt.Println("num1 + num2 =", num1+num2)
	fmt.Println("num1 - num2 =", num1-num2)
	fmt.Println("num1 * num2 =", num1*num2)
	fmt.Println("num1 / num2 =", num1/num2)
	fmt.Println("num1 % num2 =", num1%num2)
	// 참고로 파이썬에서 몫을 구해주는 "//" 연산은 go에서는 지원하지 않음

	// 문자열
	fmt.Println("\nstr1 + str2 =", str1+" "+str2) // 문자열 끼리 + 연산하면 두 문자열을 이어줌

	// 증감 연산자
	cnt1, cnt2 := 1, 10.4

	cnt1++ // 1증가
	cnt2-- // 1감소

	fmt.Println("\ncnt1++ :", cnt1)
	fmt.Println("cnt2-- :", cnt2)
	// Go에서는 이렇게 후위연산자가 가능하지만 전위연산자 (++cnt1) 은 불가능하다

	// 할당 연산자
	var a int = 3
	var num int

	num = a
	fmt.Println("\nnum = a ->", num)
	num += 3
	fmt.Println("num += 3 ->", num)
	num -= 4
	fmt.Println("num -= 4 ->", num)
	num *= 5
	fmt.Println("num *= 5 ->", num)
	num /= 2
	fmt.Println("num /= 2 ->", num)
	num %= 3
	fmt.Println("num %= 3 ->", num)

	num = 3  // 00000011
	num &= 2 // 00000011 AND 00000010 -> 00000010
	fmt.Printf("num &= 2 -> %08b, %d\n", num, num)
	num |= 5 // 00000010 OR 00000101 -> 00000111
	fmt.Printf("num |= 5 -> %08b, %d\n", num, num)
	num ^= 4 // 00000111 XOR 00000100 -> 00000011
	fmt.Printf("num ^= 4 -> %08b, %d\n", num, num)
	num &^= 2 // 00000011 NAND 00000010 -> 00000001
	fmt.Printf("num &^= 2 -> %08b, %d\n", num, num)
	num <<= 9 // 00000001 의 비트를 왼쪽으로 9칸 이동
	fmt.Printf("num <<= 9 -> %08b, %d\n", num, num)
	num >>= 8 // 1000000000 의 비트를 오른쪽으로 8칸 이동
	fmt.Printf("num <<= 9 -> %08b, %d\n", num, num)

	var b bool = true
	c := false

	// AND 연산
	fmt.Println("\n0 && 0 :", c && c)
	fmt.Println("0 && 1 :", c && b)
	fmt.Println("0 && 0 :", b && b)

	// OR 연산
	fmt.Println("0 || 0 :", c || c)
	fmt.Println("0 || 1 :", c || b)
	fmt.Println("1 || 1 :", b || b)

	// NOT 연산
	fmt.Println("!1 ", !true)
	fmt.Println("!0 ", !false)

}
