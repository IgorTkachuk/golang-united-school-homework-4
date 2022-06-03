package string_sum

import (
	"errors"
	"fmt"
	"strconv"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	clearInput := ""
	for _, s := range input {
		ss := string(s)
		if ss != " " && ss != "\t" && ss != "\n" {
			clearInput += string(ss)
		}
	}

	if len(clearInput) == 0 {
		return "", errorEmptyInput
	}

	operands := []string{""}

	j := 0
	for idx, s := range clearInput {
		if idx != 0 && (string(s) == "+" || string(s) == "-") {
			j++
			operands = append(operands, "")
		}
		operands[j] += string(s)
	}

	if len(operands) != 2 {
		return "", errorNotTwoOperands
	}

	var sum int64 = 0
	for _, o := range operands {
		n, err := strconv.ParseInt(o, 10, 32)
		if err != nil {
			return "", fmt.Errorf("error when check operand %w", err)
		}
		sum += n
	}

	output = strconv.FormatInt(sum, 10)

	return output, nil
}
