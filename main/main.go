package main

import "fmt"

func main() {

	var operator string
	var number1, number2 float64

	// Запрос первого числа у пользователя
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&number1)

	// Запрос второго числа у пользователя
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&number2)

	// Запрос оператора у пользователя
	fmt.Print("Введите оператор (+, -, * или /): ")
	fmt.Scanln(&operator)

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
