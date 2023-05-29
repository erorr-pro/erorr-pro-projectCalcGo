package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	// Чтение ввода пользователя
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	expression, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения ввода: ", err)
		return
	}

	// Вычисление
	result, err := evaluateExpression(expression)
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}

	// Результат
	fmt.Println("\nРезультат: ", result)

}

// Функция для вычисления значения выражения
func evaluateExpression(expression string) (float64, error) {
	tokens := parseExpression(expression) // Разбиваем выражение на токены
	var numberStack []float64             // Стек для хранения чисел
	var operatorStack []string            // Стек для хранения операторов

	for _, token := range tokens {
		priority := priorityOperation(token) // Определяем приоритет операции

		switch {
		case priority == 1: // Если операция имеет приоритет 1, добавляем ее в стек операторов
			operatorStack = append(operatorStack, token)

		case priority == 2: // Если операция имеет приоритет 2 (скобка закрытия), выполняем операции внутри скобок
			for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1] != "(" {
				operator := operatorStack[len(operatorStack)-1]
				operatorStack = operatorStack[:len(operatorStack)-1]

				// Выполняем вычисления с двумя последними числами и оператором
				result := calculations(numberStack[len(numberStack)-2], numberStack[len(numberStack)-1], operator)
				numberStack = numberStack[:len(numberStack)-2]
				numberStack = append(numberStack, result)
			}

			if len(operatorStack) == 0 || operatorStack[len(operatorStack)-1] != "(" {
				return 0, fmt.Errorf("некорректное выражение: несбалансированные скобки")
			}

			operatorStack = operatorStack[:len(operatorStack)-1]

		case priority == 3: // Если операция имеет приоритет 3 (сложение или вычитание), выполняем операции с предыдущими операциями любого приоритета
			for len(operatorStack) > 0 && (operatorStack[len(operatorStack)-1] == "*" || operatorStack[len(operatorStack)-1] == "/" || operatorStack[len(operatorStack)-1] == "+" || operatorStack[len(operatorStack)-1] == "-") {
				operator := operatorStack[len(operatorStack)-1]
				operatorStack = operatorStack[:len(operatorStack)-1]

				// Выполняем вычисления с двумя последними числами и оператором
				result := calculations(numberStack[len(numberStack)-2], numberStack[len(numberStack)-1], operator)
				numberStack = numberStack[:len(numberStack)-2]
				numberStack = append(numberStack, result)
			}

			operatorStack = append(operatorStack, token) // Добавляем текущую операцию в стек операторов

		case priority == 4: // Если операция имеет приоритет 4 (умножение или деление), выполняем операции с предыдущими операциями такого же приоритета
			for len(operatorStack) > 0 && (operatorStack[len(operatorStack)-1] == "*" || operatorStack[len(operatorStack)-1] == "/") {
				operator := operatorStack[len(operatorStack)-1]
				operatorStack = operatorStack[:len(operatorStack)-1]

				// Выполняем вычисления с двумя последними числами и оператором
				result := calculations(numberStack[len(numberStack)-2], numberStack[len(numberStack)-1], operator)
				numberStack = numberStack[:len(numberStack)-2]
				numberStack = append(numberStack, result)
			}

			operatorStack = append(operatorStack, token) // Добавляем текущую операцию в стек операторов

		default: // Если токен не является операцией, то он представляет число
			number, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("некорректное выражение")
			}

			numberStack = append(numberStack, number)
		}
	}

	// Выполняем оставшиеся операции в стеке операторов
	for len(operatorStack) > 0 {
		operator := operatorStack[len(operatorStack)-1]
		operatorStack = operatorStack[:len(operatorStack)-1]

		// Выполняем вычисления с двумя последними числами и оператором
		result := calculations(numberStack[len(numberStack)-2], numberStack[len(numberStack)-1], operator)
		numberStack = numberStack[:len(numberStack)-2]
		numberStack = append(numberStack, result)
	}

	return numberStack[0], nil // Возвращаем результат вычислений
}

// Функция для разбиения выражения на токены
func parseExpression(expression string) []string {

	// Проверка наличия некорректных символов в выражении
	match, _ := regexp.MatchString("[a-zA-Zа-яА-Я]+", expression)
	if match {
		fmt.Println("Ошибка: введены некорректные символы")
		log.Fatal("Произошла ошибка.")
	}

	var tokens []string
	expression = strings.ReplaceAll(expression, " ", "")
	expression = strings.ReplaceAll(expression, "(-", "(0-")
	expression = strings.ReplaceAll(expression, ")+", ")+0+")
	expression = strings.ReplaceAll(expression, "(-(", "(0-(")
	expression = strings.ReplaceAll(expression, "*-", "*(0-1)*")
	expression = strings.ReplaceAll(expression, "/-", "/(0-1)/")
	expression = strings.ReplaceAll(expression, "+-", "+(0-1)*")

	if strings.HasPrefix(expression, "-") {
		expression = "0" + expression
	}

	pattern := `[*/\-+()]|(\d+(\.\d+)?)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(expression, -1)

	for _, match := range matches {
		tokens = append(tokens, match)
	}

	return tokens
}

// Функция для выполнения элементарных вычислений в зависимости от оператора
func calculations(number1 float64, number2 float64, operator string) float64 {
	switch operator {
	case "*":
		return number1 * number2
	case "/":
		if number2 != 0 {
			return number1 / number2
		} else {
			fmt.Println("Ошибка: деление на ноль")
			return 0
		}
	case "+":
		return number1 + number2
	case "-":
		return number1 - number2
	default:
		fmt.Println("Ошибка: недопустимый оператор")
		return 0
	}
}

// priorityOperation определяет приоритет операции
func priorityOperation(token string) int {
	switch token {
	case "(":
		return 1
	case ")":
		return 2
	case "+", "-":
		return 3
	case "*", "/":
		return 4
	default:
		return 5
	}
}
