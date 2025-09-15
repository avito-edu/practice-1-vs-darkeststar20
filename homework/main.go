package main

import "fmt"

func Add(a, b float64) float64 {
	return a + b
}

func Divide(a, b float64) float64 {
	if a == 0 || b == 0 {
	    panic("На ноль делить нельзя!")
	} else {
        //fmt.Println("На ноль делить нельзя!")
        return a / b
	}
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func main() {
	var operation string
	var a, b float64
	fmt.Print("Введите число a")
	fmt.Scan(&a)
	fmt.Print("Введите число b")
	fmt.Scan(&b)
	fmt.Print("Введите название операции")
	_, err := fmt.Scan(&operation)
	if err != nil {
	    fmt.Println("Error: ", err)
	}

    fmt.Println("Вы ввели некорректное число", a, b)

	switch operation {
	case "Add":
        fmt.Println(Add(a,b))
    case "Divide":
        fmt.Println(Divide(a,b))
    case "Subtract":
        fmt.Println(Divide(a,b))
    case "Multiply":
        fmt.Println(Multiply(a,b))
    default:
        fmt.Println("Error: Неверное название операции")
    }
    defer fmt.Println("Программа завершена.")
}
