package main

import "fmt"

func main() {
	// Ініціалізація змінних
	var userinit8 uint8 = 1
	var userinit16 uint16 = 2
	var userinit64 int64 = -3
	var userautoinit = -4 // Такий варіант ініціалізації також можливий

	fmt.Println("Values: ", userinit8, userinit16, userinit64, userautoinit, "\n")

	// Короткий запис оголошення змінної
	// тільки для нових змінних
	intVar := 10

	fmt.Printf("Value = %d Type = %T\n\n", intVar, intVar)

	// Завдання.
	// 1. Вивести типи всіх змінних
	fmt.Println("Завдання 1. Вивести типи всіх змінних.")
	fmt.Printf("Тип змінної userinit8 = %T\n", userinit8)
	fmt.Printf("Тип змінної userinit16 = %T\n", userinit16)
	fmt.Printf("Тип змінної userinit64 = %T\n", userinit64)
	fmt.Printf("Тип змінної userautoinit = %T\n\n", userautoinit)

	// 2. Присвоїти змінній intVar змінні userinit16 і userautoinit. Результат вивести в консоль.
	fmt.Println("Завдання 2. Присвоїти змінній intVar змінні userinit16 та userautoinit. Результат вивести в консоль.")
	intVar = int(userinit16)
	fmt.Printf("intVar = userinit16 -> Значення = %d, Тип = %T\n", intVar, intVar)
	intVar = userautoinit
	fmt.Printf("intVar = userautoinit -> Значення = %d, Тип = %T\n", intVar, intVar)
}
