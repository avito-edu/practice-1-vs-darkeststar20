package main

import (
	"fmt"
)

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) float64 {
	if b == 0 {
		panic("На ноль делить нельзя!")
	}
	return a / b
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ошибка: ", r)
		}
		fmt.Println("Программа завершена.")
	}()

	var a, b float64
	var operation string
	var result float64

	for {
		fmt.Print("Введите число a: ")
		_, err := fmt.Scan(&a)
		if err != nil {
			fmt.Println("Ошибка ввода. Пожалуйста, введите число")
			fmt.Scan()
			continue
		}

		fmt.Print("Введите число b: ")
		_, err = fmt.Scan(&b)
		if err != nil {
			fmt.Println("Ошибка ввода. Пожалуйста, введите число")
			fmt.Scan()
			continue
		}

		fmt.Print("Введите операцию (+, -, *, /): ")
		_, err = fmt.Scan(&operation)
		if err != nil {
			fmt.Println("Ошибка ввода. Пожалуйста, введите операцию")
			fmt.Scan()
			continue
		}

		switch operation {
		case "+":
			result = Add(a, b)
		case "-":
			result = Subtract(a, b)
		case "*":
			result = Multiply(a, b)
		case "/":
			result = Divide(a, b)
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}
		default:
			fmt.Println("Неверная операция. Пожалуйста, введите +, -, * или /")
			continue
		}

		fmt.Println("Результат: ", result)
		break
	}
}
