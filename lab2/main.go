package main

import (
	"fmt"
	"lab2/utils"
)

func main() {
	var a, b, c float32 = 4.6, 2.5, 7.5
	fmt.Printf("Мінімальне серед %.1f, %.1f, %.1f = %.1f\n", a, b, c, utils.FindMin(a, b, c))
	fmt.Printf("Середнє арифметичне чисел %.1f, %.1f, %.1f = %.4f\n", a, b, c, utils.FindAvg(a, b, c))
	root, err := utils.SolveEquation(a, b)
	if err != nil {
		fmt.Println("Помилка: ", err)
	} else {
		fmt.Printf("Рішення рівняння %.1fx + %.1f = 0 -> x = %.4f", a, b, root)
	}
}
