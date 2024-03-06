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
	var dividend, divisor int

	fmt.Print("Enter the dividend: ")
	_, err := fmt.Scan(&dividend)
	if err != nil {
		fmt.Println("Error occurred while reading dividend:", err)
		return
	}

	fmt.Print("Enter the divisor: ")
	_, err = fmt.Scan(&divisor)
	if err != nil {
		fmt.Println("Error occurred while reading divisor:", err)
		return
	}

	result, err := divide(dividend, divisor)
	if err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("Result of division:", result)
	}
}
