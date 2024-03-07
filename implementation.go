package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: document this function.
// PrefixCalculate converts

func isOperand(s string) bool {
    _, err := strconv.Atoi(s)
    return err == nil
}

func isOperator(s string) bool {
    return s == "+" || s == "-" || s == "*" || s == "/"
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

    var result int
    switch operator {
    case "+":
        result = num1 + num2
    case "-":
        result = num1 - num2
    case "*":
        result = num1 * num2
    case "/":
        if num2 == 0 {
            return "", fmt.Errorf("division by zero")
        }
        result = num1 / num2
    default:
        return "", fmt.Errorf("invalid operator: %s", operator)
    }

    return strconv.Itoa(result), nil
}

func PrefixCalculate(input string) (string, error) {
	parts := strings.Fields(input)
	stack := []string{}
	for i := len(parts) - 1; i >= 0; i-- {
		if isOperand(parts[i]) {
			stack = append(stack, parts[i])
		} else if isOperator(parts[i]) {
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid prefix notation")
			}
			operand1 := stack[len(stack)-1]
			operand2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			result, error := calculate(parts[i], operand1, operand2)
			if(error != nil) {
				return "", error
			}
			stack = append(stack, result)
		} else {
			return "", fmt.Errorf("invalid element in prefix notation")
		}
	}

	if len(stack) != 1 {
		return "", fmt.Errorf("invalid prefix notation")
	}

	return stack[0], nil;
}
