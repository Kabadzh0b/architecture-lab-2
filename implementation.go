package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

// isOperand checks if the given string is a valid operand.
// It returns true if the string can be converted to an integer, false otherwise.
func isOperand(s string) bool {
    _, err := strconv.Atoi(s)
    return err == nil
}

// isOperator checks if the given string is a valid operator.
// It returns true if the string is one of the following: "+", "-", "*", "/", false otherwise.
func isOperator(s string) bool {
    return s == "+" || s == "-" || s == "*" || s == "/"
}

// calculate performs the given operation on the two operands.
// It returns the result of the operation as a string and an error if the operation is invalid.
func calculate(operator, operand1, operand2 string) (string, error) {
    num1, err := strconv.ParseFloat(operand1, 64)
    if err != nil {
        return "", fmt.Errorf("invalid operand: %s", operand1)
    }

    num2, err := strconv.ParseFloat(operand2, 64)
    if err != nil {
        return "", fmt.Errorf("invalid operand: %s", operand2)
    }

    var result float64
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

    // Check if the result is an integer
    if result == float64(int(result)) {
        // If it is, format it without decimal places
        return fmt.Sprintf("%.0f", result), nil
    }

    // Otherwise, format it with six digits after the decimal point
    return fmt.Sprintf("%.6f", result), nil
}

// PrefixCalculate computes the result of the given prefix notation expression.
// It returns the result as a string and an error if the expression is invalid.
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
