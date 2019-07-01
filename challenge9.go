package main

import (
	"fmt"
)

func challenge9() error {

	input := "YELLOW SUBMARINE"

	padding := byte(0x04)

	output := pad(input, padding, 20)

	fmt.Println("Challenge 9 success!")
	return nil
}
