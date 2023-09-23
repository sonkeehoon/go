package main

import (
	"fmt"
	"unsafe"
	// unsafe는 자료형이 몇바이트를 사용하는지 확인할수 있는 패키지
)

func main() {
	var a byte
	var b int        // int8, int16, int32, int64
	var c uint       // uint8(=byte), uint16, uint32, uint64, uintptr
	var d float64    // float은 뒤에 32나 64를 반드시
	var e complex128 // complex는 뒤에 64나 128을 반드시
	var f string
	var g rune // (=int32)

	fmt.Println(unsafe.Sizeof(a)) // byte 타입은 1바이트 크기
	fmt.Println(unsafe.Sizeof(b)) // int 타입은 8바이트 크기
	fmt.Println(unsafe.Sizeof(c)) // uint 타입은 8바이트 크기
	fmt.Println(unsafe.Sizeof(d)) // float64 타입은 8바이트 크기
	fmt.Println(unsafe.Sizeof(e)) // complex128 타입은 16바이트 크기
	fmt.Println(unsafe.Sizeof(f)) // string 타입은 16바이트 크기
	fmt.Println(unsafe.Sizeof(g)) // rune 타입은 4바이트 크기

}
