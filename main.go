package main

import (
	"fmt"
	"strings"
)

func BraketOf(expression string) (string, error) {
	indexOpen := strings.Index(expression, "(")
	indexClose := strings.Index(expression, ")")

	for indexOpen > indexClose {
		if indexClose == -1 {
			return "", fmt.Errorf("400 Скобка должна быть закрыта")
		}
		if indexClose < indexOpen {
			// Выполнение выражения без скобок
			return expression, nil
		}
		BraketOf(expression[:indexOpen])
		indexOpen = strings.Index(expression, "(")
		indexClose = strings.Index(expression, ")")

	}

	return expression, nil

}

func Calc(expression string) (float64, error) {
	expression = strings.Replace(expression, " ", "", -1)
	index := strings.Index(expression, "(")
	for index != -1 {
		BraketOf(expression[index:])
	}
	var return_float float64
	return return_float, nil
}
func main() {
	Calc("()")
}
