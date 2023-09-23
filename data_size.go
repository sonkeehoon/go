package main

import (
	"fmt"
	"unsafe"
	// unsafe는 자료형이 몇바이트를 사용하는지 확인할수 있는 패키지
)

func main() {
	var a byte
	var b int

	fmt.Println(unsafe.Sizeof(a)) // byte 타입은 1바이트 크기
	fmt.Println(unsafe.Sizeof(b)) // int 타입은 8바이트 크기

}
