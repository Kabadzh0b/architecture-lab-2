package lab2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func isOperand(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/" || s == "^"
}

func calculate(operator, operand1, operand2 string) (string, error) {
	num1, err := strconv.Atoi(operand1)
	if err != nil {
			return "", fmt.Errorf("invalid operand: %s", operand1)
	}

	num2, err := strconv.Atoi(operand2)
	if err != nil {
			return "", fmt.Errorf("invalid operand: %s", operand2)
	}

	var result float64
	switch operator {
	case "+":
			result = float64(num1) + float64(num2)
	case "-":
			result = float64(num1) - float64(num2)
	case "*":
			result = float64(num1) * float64(num2)
	case "^":
			result = math.Pow(float64(num1), float64(num2))
	case "/":
			if num2 == 0 {
					return "", fmt.Errorf("division by zero")
			}
			result = float64(num1) / float64(num2)
	default:
			return "", fmt.Errorf("invalid operator: %s", operator)
	}

	return strconv.Itoa(int(result)), nil
}

func PrefixCalculate(input string) (string, error) {
	parts := strings.Fields(input)
	stack := []string{}
	for i := len(parts) - 1; i >= 0; i-- {
		if isOperand(parts[i]) {
			stack = append(stack, parts[i])
		} else if isOperator(parts[i]) {
			if len(stack) < 2 {
				return "invalid prefix notation: insufficient operands", fmt.Errorf("invalid prefix notation: not enough operands")
			}
			operand1 := stack[len(stack)-1]
			operand2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			result, err := calculate(parts[i], operand1, operand2)
			if err != nil {
				return "invalid prefix notation: insufficient operands", err
			}
			stack = append(stack, result)
		} else {
			return "invalid prefix notation: insufficient operands", fmt.Errorf("invalid prefix notation: invalid element")
		}
	}

	if len(stack) != 1 {
		return "invalid prefix notation: insufficient operands", fmt.Errorf("invalid prefix notation: insufficient operands")
	}

	return stack[0], nil
}

