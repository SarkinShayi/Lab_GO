package main

//Импорт нескольких пакетов
import (
	"fmt"
	"math"
)

func main() {
	var defaultFloat float32
	var defaultDouble float64 = 5.5

	fmt.Println("defaultfloat       = ", defaultFloat)
	fmt.Printf("defaultDouble (%T) = %f\n\n", defaultDouble, defaultDouble)

	fmt.Println("MAX float32        = ", math.MaxFloat32)
	fmt.Println("MIN float32        = ", math.SmallestNonzeroFloat32, "\n")

	fmt.Println("MAX float64        = ", math.MaxFloat64)
	fmt.Println("MIN float64        = ", math.SmallestNonzeroFloat64, "\n\n")

	// Завдання.
	// 1. Створіть змінні різних типів, використовуючи короткий запис та ініціалізацію за замовчуванням. Результат вивести в консоль
	intVal := 245
	var defaultInt int
	strVal := "Text"
	var defaultString string
	boolVal := true
	var defaultBool bool
	byteVal := byte(16)
	var defaultByte byte
	runeVal := 'X'
	var defaultRune rune

	fmt.Println("Завдання 1. Створіть змінні різних типів, використовуючи короткий запис та ініціалізацію за замовчуванням. Результат вивести в консоль.")
	fmt.Printf("intVal        (%T) = %d\n", intVal, intVal)
	fmt.Printf("defaultInt    (%T) = %d\n", defaultInt, defaultInt)
	fmt.Printf("strVal        (%T) = %q\n", strVal, strVal)
	fmt.Printf("defaultString (%T) = %q\n", defaultString, defaultString)
	fmt.Printf("boolVal       (%T) = %t\n", boolVal, boolVal)
	fmt.Printf("defaultBool   (%T) = %t\n", defaultBool, defaultBool)
	fmt.Printf("byteVal       (%T) = %d\n", byteVal, byteVal)
	fmt.Printf("defaultByte   (%T) = %d\n", defaultByte, defaultByte)
	fmt.Printf("runeVal       (%T) = %d\n", runeVal, runeVal)
	fmt.Printf("defaultRune   (%T) = %d", defaultRune, defaultRune)
}
