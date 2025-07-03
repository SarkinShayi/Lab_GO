package main

import "fmt"

func main() {
	variable8 := int8(127)
	variable16 := int16(16383)

	fmt.Println("Приведення типів\n")

	fmt.Printf("variable8         = %-5d = %.16b\n", variable8, variable8)
	fmt.Printf("variable16        = %-5d = %.16b\n", variable16, variable16)
	fmt.Printf("uint16(variable8) = %-5d = %.16b\n", uint16(variable8), uint16(variable8))
	fmt.Printf("uint8(variable16) = %-5d = %.16b\n\n", uint8(variable16), uint8(variable16))

	// Завдання.
	// 1. Створіть 2 змінні різних типів. Виконайте арифметичні операції. Результат вивести в консоль
	var a int = 5
	var pi float32 = 3.1415
	fmt.Println("Завдання 1. Створіть 2 змінні різних типів. Виконайте арифметичні операції. Результат вивести в консоль.")
	fmt.Printf("a = %d (тип %T)\n", a, a)
	fmt.Printf("pi = %.4f (тип %T)\n\n", pi, pi)

	fmt.Printf("Сума:     a + pi = %d\n", a+int(pi))
	fmt.Printf("Різниця:  a - pi = %d\n", a-int(pi))
	fmt.Printf("Множення: a * pi = %d\n", a*int(pi))
	fmt.Printf("Ділення:  a / pi = %d\n", a/int(pi))
}
