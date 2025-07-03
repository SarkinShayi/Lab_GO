package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Синоніми цілих типів\n")

	fmt.Println("byte - int8")
	fmt.Println("rune - int32")
	fmt.Println("int  - int32 або int64, в залежності від платформи")
	fmt.Println("uint - uint32 або uint64, в залежності від платформи\n")

	// Завдання.
	// 1. Визначити розрядність платформи
	var i int
	fmt.Println("Завдання 1. Визначити розрядність платформи.")
	fmt.Printf("Розмір int: %d байт\n", unsafe.Sizeof(i))
	if unsafe.Sizeof(i) == 4 {
		fmt.Println("Розрядність платформи: 32-біт")
	} else if unsafe.Sizeof(i) == 8 {
		fmt.Println("Розрядність платформи: 64-біт")
	} else {
		fmt.Println("Нестандартна розрядність платформи")
	}
}
