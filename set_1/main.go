package main

import (
	"fmt"
)

func main() {
	var err error
	err = challenge1()
	if err != nil {
		fmt.Println("Challenge 1 failed.")
		fmt.Println(err)
	}
	err = challenge2()
	if err != nil {
		fmt.Println("Challenge 2 failed.")
		fmt.Println(err)
	}
	err = challenge3()
	if err != nil {
		fmt.Println("Challenge 3 failed.")
		fmt.Println(err)
	}
}
