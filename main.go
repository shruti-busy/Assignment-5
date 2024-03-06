package main

import (
	"fmt"
)

func divide(x, y int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()

	if y == 0 {
		panic("division by zero")
	}

	result = x / y
	return result, nil
}

func main() {
	dividend := 10
	divisor := 0

	result, err := divide(dividend, divisor)
	if err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("Result of division:", result)
	}
}
