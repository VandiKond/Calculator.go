package main

import "testing"

func TestCalc(t *testing.T) {
	// Тесты для корректных выражений
	testCases := []struct {
		expression string
		expected   float64
	}{
		{"(1 + 2) * 3 / 4", 2.25},
	}

	for _, testCase := range testCases {
		t.Run(testCase.expression, func(t *testing.T) {
			result, err := Calc(testCase.expression)
			if err != nil {
				t.Errorf("Calc(%s) error: %v", testCase.expression, err)
			} else if result != testCase.expected {
				t.Errorf("Calc(%s) = %v, want %v", testCase.expression, result, testCase.expected)
			}
		})
	}
}
