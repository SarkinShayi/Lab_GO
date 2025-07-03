package main

import "fmt"

func main() {
	var chartype int8 = 'R'

	fmt.Printf("Code '%c' - %d\n\n", chartype, chartype)

	// Завдання.
	// 1. Вивести українську літеру 'Ї'
	fmt.Println("Завдання 1. Вивести українську літеру 'Ї'.")
	var letter rune = '\u0407'
	fmt.Printf("Літера: %c", letter)

	// 2. Пояснити призначення типу "rune"
}
