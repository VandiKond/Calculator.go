package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Все разрешенные операторы
var operators = []string{"*", "+", "/", "-"}

// Структурадля операции
type Operation struct {
	// Первое число
	num1 float64
	// Тип операции
	symvol string
	// Второе число
	num2 float64
}

// Метод для выполнения операции
func (Op Operation) ParseOper() (float64, error) {
	// Предваритеьно создаем конечное число
	var num float64
	// Проходимся по символам
	switch Op.symvol {
	// В случае *, +, - возвращаем результат этих дествий
	case "*":
		num = Op.num1 * Op.num2
		break
	case "+":
		num = Op.num1 + Op.num2
		break
	case "-":
		num = Op.num1 - Op.num2
		break
	// В случае деления проверяем второе число на то, что оно 0. Если нет то возвращаем результат деления
	case "/":
		if Op.num2 == 0 {
			return 0, fmt.Errorf("Делить на 0 нельзя")
		}
		num = Op.num1 / Op.num2
		break
	// В ином случае вызываем ошибку
	default:
		return 0, fmt.Errorf("400 Неизвестный знак : [%s]", Op.symvol)
	}
	// Возвращаем результат
	return num, nil
}

// Метод для переделавания из примера строку
func (Op Operation) FormatToString() string {
	// Возвращаем пример ввиде стороки
	return strconv.FormatFloat(Op.num1, 'f', -1, 64) + Op.symvol + strconv.FormatFloat(Op.num2, 'f', -1, 64)
}

// Функция для получения индекса по условию.
func findIndex(str string, findType func(int, int) bool, stand_result int, getIndex func(str, operator string) int) (int, string) {
	// Задаем стандартный результат
	ResultIndex := stand_result
	var operatorResult string
	// Проходимся по операциям
	for _, operator := range operators {
		// Получаем индекс
		index := getIndex(str, operator)

		// Проверяем существование индекса
		if index == -1 {
			continue
		}

		// Проверка на унарный минус
		if operator == "-" && index == 0 {
			// Получаем слудуйщий оператор после -
			index, operator = findIndex(str[1:], findType, stand_result, getIndex)
			index++
		}

		// Проверяем индекс по условию и в случае успеха задаем данными переменных результатами
		if findType(index, ResultIndex) {
			ResultIndex = index
			operatorResult = operator
		}

	}

	// Возвращаем индекс
	return ResultIndex, operatorResult
}

// Функция для операций только из + и -
func OrderOperations(expression string) (string, error) {
	// Получение индекса первого оператора
	index, operator := findIndex(expression, func(i1, i2 int) bool { return i1 < i2 }, len(expression), strings.Index)

	// Конвертируем в число все что до оператора. В случае ошибки возвращаем ее
	num1, err := strconv.ParseFloat(expression[:index], 64)
	if err != nil {
		return expression, fmt.Errorf("Ошибка обработки: [%s] не является числом", expression[:index])
	}

	// Получаем индекс второго оператор. Конвертируем в число его. В случае ошибки возвращаем ее
	expressionTilEnd := expression[index+1:]
	indexOfEnd, _ := findIndex(expressionTilEnd, func(i1, i2 int) bool { return i1 < i2 }, len(expressionTilEnd), strings.Index)
	num2, err := strconv.ParseFloat(expressionTilEnd[:indexOfEnd], 64)
	if err != nil {
		return expression, fmt.Errorf("Ошибка обработки: [%s] не является числом", expressionTilEnd[:indexOfEnd])
	}

	// СОздаем данные об операции
	oper := Operation{num1: num1, symvol: operator, num2: num2}

	// Выполняем операцию. в случае ошибки возвращаем ее
	result, err := oper.ParseOper()
	if err != nil {
		return expression, fmt.Errorf("Ошибка выполнения примера [%s]: %v", oper.FormatToString(), err)
	}
	// Заменяем выполненую операцию
	expression = strings.Replace(expression, oper.FormatToString(), strconv.FormatFloat(result, 'f', -1, 64), 1)

	// Проверяем является ли результат числом
	_, ok := strconv.ParseFloat(expression, 64)
	if ok != nil {
		// В случае если результат не число, то продолжаем операцию
		expression, err = OrderOperations(expression)
		if err != nil {
			return expression, fmt.Errorf("Ошибка обработки: [%s] не является числом", expression)
		}
	}

	// Возвращаем результат операции
	return expression, nil
}

// Функция для приоретезации порядка в выражение без скобок
func ManageOrder(expression string) (string, error) {
	// Проверяем не является ли выражение просто числом
	_, err := strconv.ParseFloat(expression, 64)
	if err == nil {
		return expression, nil
	}

	// Полуаем индекс умножение и делания
	indexMul := strings.Index(expression, "*")
	indexDiv := strings.Index(expression, "/")

	// В случае их отстутсвия заменяем их на длинну строки
	if indexDiv == -1 {
		indexDiv = len(expression)
	}
	if indexMul == -1 {
		indexMul = len(expression)
	}

	// В случае отсутствия их обоих идем по порядку
	if indexDiv == indexMul {
		expression, err := OrderOperations(expression)
		if err != nil {
			return expression, fmt.Errorf("Ошибка последовательной операции [%s]: %v", expression, err)
		}
		return expression, nil
	}

	// Создаем пустой индекс и тип
	index := -1
	operatType := ""

	// В случае когда умножение первое задаем индекс на индекс умножения и тип операции на умножения
	if indexMul < indexDiv && indexMul != len(expression) {
		index = indexMul
		operatType = "*"
	} else /*Иначе делаеи все для деления*/ {
		index = indexDiv
		operatType = "/"
	}
	// Получаем выражение до и после операции
	expressionBe4 := expression[:index]
	expressionAfter := expression[index+1:]

	// Получаем индексы ближайших операций
	indexBe4, _ := findIndex(expressionBe4, func(i1, i2 int) bool { return i1 > i2 }, -1, strings.LastIndex)
	indexAfter, _ := findIndex(expressionAfter, func(i1, i2 int) bool { return i1 < i2 }, len(expressionAfter), strings.Index)

	// Получаем цифры для операции. В случае ошибки возвращаем ошибку
	num1, err1 := strconv.ParseFloat(expressionBe4[indexBe4+1:], 64)
	num2, err2 := strconv.ParseFloat(expressionAfter[:indexAfter], 64)
	if err1 != nil {
		return expression, fmt.Errorf("Ошибка обработки: [%s] не является числом", expressionBe4[indexBe4+1:])
	}
	if err2 != nil {
		return expression, fmt.Errorf("Ошибка обработки: [%s] не является числом", expressionAfter[:indexAfter])
	}

	// Создаем операцию
	oper := Operation{num1: num1, symvol: operatType, num2: num2}

	// Выполнение операции. в случае ошибки возвращаем ее
	result, err := oper.ParseOper()
	if err != nil {
		return expression, fmt.Errorf("Ошибка выполнения примера [%s]: %v", oper.FormatToString(), err)
	}

	// Заменяем выполненую операцию
	expression = strings.Replace(expression, oper.FormatToString(), strconv.FormatFloat(result, 'f', -1, 64), 1)

	// Проверяем является ли результат числом
	_, ok := strconv.ParseFloat(expression, 64)
	if ok != nil {
		// В случае если результат не число, то продолжаем операцию
		expression, err = ManageOrder(expression)
		if err != nil {
			return expression, fmt.Errorf("Ошибка выполнения примера [%s]: %v", expression, err)
		}
	}

	// Возвращаем результат операции
	return expression, nil
}

// Функция для приоретезации скобок
func BraketOf(expression string) (string, error) {

	// Поиск первой ( и )
	indexOpen := strings.Index(expression, "(")
	indexClose := strings.Index(expression, ")")

	// Повторяем пока открывающиеся скобки не станут после закрывающихся или не закончатся
	for indexClose > indexOpen || indexOpen != -1 {

		// В случчае отсутствия закрывающися скобок при наличии открых отправляем ошибку
		if indexOpen != indexClose && indexClose == -1 {
			return expression, fmt.Errorf("Скобка должна быть закрыта")
		}

		// В случае Последней открытой скобки
		if indexClose < indexOpen || indexOpen == -1 {

			// Выполняем пример без скобок, в случае ошибки возвращаем ошибку
			managedexeption, err := ManageOrder(expression[:indexClose])
			if err != nil {
				return expression, fmt.Errorf("Ошибка выполнения примера [%s]: %v", expression, err)
			}
			// Заменяем выражение на исход примера и возвращаем оставшееся выражение
			expression = strings.Replace(expression, expression[:indexClose+1], managedexeption, 1)
			return expression, nil
		}

		// С помощью рекурсии идем внутрь следуших скобок, в случае ошибки возвращаем ошибку
		BreketOfEx, err := BraketOf(expression[indexOpen+1:])
		if err != nil {
			return expression, fmt.Errorf("Ошибка внутри выражения [%s]: %v", expression[indexOpen+1:], err)
		}

		// Заменяем строку на результат убиранее скобок и идем по циклу дальше
		expression = strings.Replace(expression, expression[indexOpen:], BreketOfEx, 1)
		indexOpen = strings.Index(expression, "(")
		indexClose = strings.Index(expression, ")")
	}

	// Возвращаем выражение
	return expression, nil

}

// Функция калькулятора
func Calc(expression string) (float64, error) {
	// Создание ошибки зарание как пустой
	var err error = nil

	// Убираем пробелы
	expression = strings.Replace(expression, " ", "", -1)

	// Ищем первую скобку и проверяем ее количество
	index := strings.Index(expression, "(")
	if index != -1 {

		// Вызываем функцию для выполнения операций внутри скобок. В случае ошибки возврашаем ошибку
		expression, err = BraketOf(expression)
		if err != nil {
			return float64(0), err
		}
	}

	// Выполняем операцию с оставшийся строкой без скобок. В случае ошибки возврщаем ошибку
	expression, err = ManageOrder(expression)
	if err != nil {
		return float64(0), fmt.Errorf("Ошибка выполнения примера [%s]: %v", expression, err)
	}

	return strconv.ParseFloat(expression, 64)
}
