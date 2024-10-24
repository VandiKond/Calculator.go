package main

import "testing"

func TestCalc(t *testing.T) {
	// Тесты для корректных выражений
	testCases := []struct {
		expression string
		expected   float64
	}{
		{"1+1", 2},
		{"3+3*6", 21},
		{"1+8/2*4", 17},
		{"(1+1) *2", 4},
		{"((1+4) * (1+2) +10) *4", 100},
		{"(4+3+2)/(1+2) * 10 / 3", 10},
		{"(70/7) * 10 /((3+2) * (3+7)) -2", 0},
		{"((7+1) / (2+2) * 4) / 8 * (32 - ((4+12)*2)) -1", -1},
		{"-1", -1},
		{"+5", 5},
		{"5+5+5+5+5", 25},
		{"(1)", 1},
		{"(1+2*(10) + 10)", 31},
		{"((1+2)*(5*(7+3) - 70 / (3+4) * (1+2)) - (8-1)) + (10 * (5-1 * (2+3)))", 53},
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
