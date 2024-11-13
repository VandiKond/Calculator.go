package calc

import (
	"fmt"
	"strconv"
	"strings"
)

// All allowed operations
var operators = []string{"*", "+", "/", "-"}

// Operation struct
type Operation struct {
	// Thirst number
	num1 float64
	// Operation type (one of operations)
	symbol string
	// Second number
	num2 float64
}

// Method to do operation
func (Op Operation) ParseOpr() (float64, error) {
	// Making the result num
	var num float64
	// Switching operation types
	switch Op.symbol {
	// In case of *, +, - Returning the operation result
	case "*":
		num = Op.num1 * Op.num2
	case "+":
		num = Op.num1 + Op.num2
	case "-":
		num = Op.num1 - Op.num2
	// If it is dividing checking that the number is not zero 0. And making the operation
	case "/":
		if Op.num2 == 0 {
			return 0, fmt.Errorf("dividing by zero is not allowed")
		}
		num = Op.num1 / Op.num2
	// Default returning an error
	default:
		return 0, fmt.Errorf("unknown symbol: [%s]", Op.symbol)
	}
	// returning the result
	return num, nil
}

// Method to create a string of the operation
func (Op Operation) FormatToString() string {
	return strconv.FormatFloat(Op.num1, 'f', -1, 64) + Op.symbol + strconv.FormatFloat(Op.num2, 'f', -1, 64)
}

// Getting operator index by findType and getIndex
func findIndex(str string, findType func(int, int) bool, stand_result int, getIndex func(str, operator string) int) (int, string) {
	// Setting the standard result
	ResultIndex := stand_result
	var operatorResult string
	// Checking all operators
	for _, operator := range operators {
		// Getting the index
		index := getIndex(str, operator)

		// Checking the index existence
		if index == -1 {
			continue
		}

		if operator == "-" && index == 0 {
			// Getting the next operator after -
			index, operator = findIndex(str[1:], findType, stand_result, getIndex)
			index++
		}

		if findType(index, ResultIndex) {
			ResultIndex = index
			operatorResult = operator
		}

	}

	// return the result
	return ResultIndex, operatorResult
}

// Order operation completing
func OrderOperations(expression string) (string, error) {
	// Getting the index
	index, operator := findIndex(expression, func(i1, i2 int) bool { return i1 < i2 }, len(expression), strings.Index)

	// Tying to convert the result to a number
	num1, err := strconv.ParseFloat(expression[:index], 64)
	if err != nil {
		return expression, fmt.Errorf("num parsing error: [%s] is not a number", expression[:index])
	}

	// Getting the index of the second operator
	expressionTilEnd := expression[index+1:]
	indexOfEnd, _ := findIndex(expressionTilEnd, func(i1, i2 int) bool { return i1 < i2 }, len(expressionTilEnd), strings.Index)
	num2, err := strconv.ParseFloat(expressionTilEnd[:indexOfEnd], 64)
	if err != nil {
		return expression, fmt.Errorf("num parsing error: [%s] is not a number", expressionTilEnd[:indexOfEnd])
	}

	// Creating the operation data
	opr := Operation{num1: num1, symbol: operator, num2: num2}

	// Doing the operation
	result, err := opr.ParseOpr()
	if err != nil {
		return expression, fmt.Errorf("error doing operation [%s]: %w", opr.FormatToString(), err)
	}
	// Replacing the operation with the result
	expression = strings.Replace(expression, opr.FormatToString(), strconv.FormatFloat(result, 'f', -1, 64), 1)

	// Checking is the result a number
	_, ok := strconv.ParseFloat(expression, 64)
	if ok != nil {
		// Continue the order oration
		expression, err = OrderOperations(expression)
		if err != nil {
			return expression, fmt.Errorf("error completing order operation [%s]: %w", expression, err)
		}
	}

	return expression, nil
}

// Manages the operation order without brackets
func ManageOrder(expression string) (string, error) {
	// Checking is the result a number
	_, err := strconv.ParseFloat(expression, 64)
	if err == nil {
		return expression, nil
	}

	// Getting the index of the dividing and multiplying
	indexMul := strings.Index(expression, "*")
	indexDiv := strings.Index(expression, "/")

	// Is they don't exist replacing with the string length
	if indexDiv == -1 {
		indexDiv = len(expression)
	}
	if indexMul == -1 {
		indexMul = len(expression)
	}

	// If they are the same (they don't exist)
	if indexDiv == indexMul {
		expression, err := OrderOperations(expression)
		if err != nil {
			return expression, fmt.Errorf("error completing order operation [%s]: %w", expression, err)
		}
		return expression, nil
	}

	// Creating empty index and type
	index := -1
	oprType := ""

	// Setting the type and index
	if indexMul < indexDiv && indexMul != len(expression) {
		index = indexMul
		oprType = "*"
	} else {
		index = indexDiv
		oprType = "/"
	}

	// getting the expression before and after
	expressionBe4 := expression[:index]
	expressionAfter := expression[index+1:]

	// Getting the nearest operation index
	indexBe4, _ := findIndex(expressionBe4, func(i1, i2 int) bool { return i1 > i2 }, -1, strings.LastIndex)
	indexAfter, _ := findIndex(expressionAfter, func(i1, i2 int) bool { return i1 < i2 }, len(expressionAfter), strings.Index)

	// Getting the numbers for the operation
	num1, err1 := strconv.ParseFloat(expressionBe4[indexBe4+1:], 64)
	num2, err2 := strconv.ParseFloat(expressionAfter[:indexAfter], 64)
	if err1 != nil {
		return expression, fmt.Errorf("num parsing error: [%s] is not a number", expressionBe4[indexBe4+1:])
	}
	if err2 != nil {
		return expression, fmt.Errorf("num parsing error: [%s] is not a number", expressionAfter[:indexAfter])
	}

	// Creating the operation
	opr := Operation{num1: num1, symbol: oprType, num2: num2}

	// Doing the operation
	result, err := opr.ParseOpr()
	if err != nil {
		return expression, fmt.Errorf("error doing operation [%s]: %w", opr.FormatToString(), err)
	}

	// Replacing the operation result
	expression = strings.Replace(expression, opr.FormatToString(), strconv.FormatFloat(result, 'f', -1, 64), 1)

	// Checking that the operation is not a number
	_, ok := strconv.ParseFloat(expression, 64)
	if ok != nil {
		// Continuing the operation
		expression, err = ManageOrder(expression)
		if err != nil {
			return expression, fmt.Errorf("error completing the expression [%s]: %w", expression, err)
		}
	}

	return expression, nil
}

// Gets rid of brackets
func BracketOf(expression string) (string, error) {

	// Find first brackets
	indexOpen := strings.Index(expression, "(")
	indexClose := strings.Index(expression, ")")

	for indexClose > indexOpen || indexOpen != -1 {

		// If close bracket not exists
		if indexOpen != indexClose && indexClose == -1 {
			return expression, fmt.Errorf("bracket should be closed")
		}

		if indexClose < indexOpen || indexOpen == -1 {
			// Doing the operation without brackets
			managedExpression, err := ManageOrder(expression[:indexClose])
			if err != nil {
				return expression, fmt.Errorf("error completing the expression [%s]: %w", expression, err)
			}
			// Replacing the result
			expression = strings.Replace(expression, expression[:indexClose+1], managedExpression, 1)
			return expression, nil
		}

		// Going into the next bracket
		BracketOfEx, err := BracketOf(expression[indexOpen+1:])
		if err != nil {
			return expression, fmt.Errorf("error in getting rid of brackets [%s]: %w", expression[indexOpen+1:], err)
		}

		// Replacing the sting
		expression = strings.Replace(expression, expression[indexOpen:], BracketOfEx, 1)
		indexOpen = strings.Index(expression, "(")
		indexClose = strings.Index(expression, ")")
	}

	return expression, nil

}

// Calculator
func Calc(expression string) (float64, error) {
	// Creating empty error
	var err error = nil

	// Deleting spaces
	expression = strings.Replace(expression, " ", "", -1)

	// Finding the first bracket
	index := strings.Index(expression, "(")
	if index != -1 {

		// Getting rid of brackets
		expression, err = BracketOf(expression)
		if err != nil {
			return float64(0), fmt.Errorf("error in getting rid of brackets [%s]: %w", expression, err)
		}
	}

	// Managing the operation without the brackets
	expression, err = ManageOrder(expression)
	if err != nil {
		return float64(0), fmt.Errorf("error completing the expression [%s]: %w", expression, err)
	}

	return strconv.ParseFloat(expression, 64)
}
