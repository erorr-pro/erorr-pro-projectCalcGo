package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Создание нового сканера для чтения ввода пользователя
	scanner := bufio.NewScanner(os.Stdin)

	var operator string
	var number1, number2 float64

	// Запрос первого числа у пользователя
	fmt.Print("Введите первое число: ")
	for {
		if scanner.Scan() {
			input := scanner.Text()
			if num, err := strconv.ParseFloat(input, 64); err == nil {
				number1 = num
				break
			} else {
				fmt.Println("Ошибка: введите корректное число")
				fmt.Print("Введите первое число: ")
			}
		}
	}

	// Запрос второго числа у пользователя
	fmt.Print("Введите второе число: ")
	for {
		if scanner.Scan() {
			input := scanner.Text()
			if num, err := strconv.ParseFloat(input, 64); err == nil {
				number2 = num
				break
			} else {
				fmt.Println("Ошибка: введите корректное число")
				fmt.Print("Введите второе число: ")
			}
		}
	}

	// Запрос оператора у пользователя
	fmt.Print("Введите оператор (+, -, * или /): ")
	for {
		if scanner.Scan() {
			operator = scanner.Text()
			if isValidOperator(operator) {
				break
			} else {
				fmt.Println("Ошибка: некорректный оператор")
				fmt.Print("Введите оператор (+, -, * или /): ")
			}
		}
	}

	// Вызов функции calculations для выполнения операции
	result := calculations(number1, number2, operator)

	// Вывод результата
	fmt.Println("Результат:", result)
}

// Функция calculations выполняет операцию над числами
func calculations(number1, number2 float64, operator string) float64 {

	var result float64

	// Используем оператор switch для выполнения нужной операции
	switch operator {
	case "+":
		result = number1 + number2
	case "-":
		result = number1 - number2
	case "*":
		result = number1 * number2
	case "/":
		// Проверяем, что делитель не равен нулю
		if number2 != 0 {
			result = number1 / number2
		} else {
			// Выводим сообщение об ошибке при делении на ноль
			fmt.Println("Ошибка: деление на ноль")
		}
	default:
		// Выводим сообщение об ошибке при некорректном операторе
		fmt.Println("Ошибка: некорректный оператор")
	}

	return result

}

// Функция проверяет, является ли оператор корректным (+, -, * или /)
func isValidOperator(operator string) bool {
	return operator == "+" || operator == "-" || operator == "*" || operator == "/"
}
