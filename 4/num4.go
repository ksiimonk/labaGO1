package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Ввод операции
	fmt.Print("Введите операцию: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)

	// Разделение ввода на части
	parts := strings.Fields(input)

	// Преобразование строк в целые числа
	num1, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}
	num2, err := strconv.Atoi(parts[2])
	if err != nil {
		log.Fatal(err)
	}

	// Выполнение операции в зависимости от знака
	switch parts[1] {
	case "+":
		fmt.Printf("Результат: %d\n", num1+num2)
	case "-":
		fmt.Printf("Результат: %d\n", num1-num2)
	case "*":
		fmt.Printf("Результат: %d\n", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Ошибка: Деление на ноль")
		} else {
			fmt.Printf("Результат: %d\n", num1/num2)
		}
	default:
		fmt.Println("Ошибка: Неизвестный оператор. Используйте +, -, * или /.")
	}
}
