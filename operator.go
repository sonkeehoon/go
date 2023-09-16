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

	// 논리 연산자
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

	// 관계 연산자
	fmt.Println("\n13 == 13 :", 13 == 13)
	fmt.Println("13 == 23 :", 13 == 23)
	fmt.Println("13 != 13 :", 13 != 13)
	fmt.Println("3 != 5 :", 3 != 5)
	fmt.Println("0 < 1 :", 0 < 1)
	fmt.Println("0 > 1 :", 0 > 1)
	fmt.Println("0 >= 1 :", 0 >= 1)
	fmt.Println("0 <= 1 :", 0 <= 1)

	// 비트 연산자
	num11 := 15        // 00001111
	var num12 int = 20 // 00010100

	fmt.Printf("\nnum11 & num12 : %08b, %d\n", num11&num12, num11&num12)  // AND
	fmt.Printf("num11 | num12 : %08b, %d\n", num11|num12, num11|num12)    // OR
	fmt.Printf("num11 ^ num12 : %08b, %d\n", num11^num12, num11^num12)    // XOR
	fmt.Printf("num11 &^ num12 : %08b, %d\n", num11&^num12, num11&^num12) // NAND

	fmt.Printf("num11 << 4 : %08b, %d\n", num11<<4, num11<<4) // 00001111비트를 왼쪽으로 4칸 이동
	fmt.Printf("num12 >> 4 : %08b, %d\n", num12>>2, num12>>2) // 00001111비트를 오른쪽으로 2칸 이동

	// 채널 연산자 : Go에만 있는 개념
	// 채널은 고루틴(goroutine)끼리 데이터를 주고받고 실행의 흐름을 제어하는 기능
	// <-는 채널의 수신을 연산합니다. 채널에 값을 보내거나 불러옵니다
	ch := make(chan int) // 정수형 채널 생성

	go func() {
		ch <- 10
	}() // 채널에 10을 보냄

	result := <-ch // 채널로부터 10을 전달받음
	fmt.Printf("\n%d\n", result)

	// 포인터 연산자
	var num3 int = 5
	pnum := &num3

	fmt.Println("\nnum3 :", num3) // num3 값
	fmt.Println("pnum :", pnum)   // num3의 메모리 주소
	fmt.Println("pnum :", *pnum)  // num3의 메모리 주소에 할당되어있는 값에 접근

	*pnum++                      // num3의 메모리 주소에 할당돼있는 값에 접근후 1을 더함
	fmt.Println("num3 :", num3)  // 기존 메모리 주소에 할당되어있는 값에서 1 더한값을 출력
	fmt.Println("pnum :", *pnum) // 기존 메모리 주소에 할당되어있는 값에서 1 더한값을 출력

}
