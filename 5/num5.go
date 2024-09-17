package main

// Импорт библиотек
import (
	"fmt"
)

func main() {
	// Ввод чисел
	fmt.Print("Введите  числа: ")
	var num1 float64
	var num2 float64
	fmt.Scan(&num1)
	fmt.Scan(&num2)

	// Выполнение  и вывод операций
	fmt.Printf("Результат: %.2f\n", num1+num2)
	fmt.Printf("Результат: %.2f\n", num1-num2)
}
