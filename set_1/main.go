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
	err = challenge4()
	if err != nil {
		fmt.Println("Challenge 4 failed.")
		fmt.Println(err)
	}
	err = challenge5()
	if err != nil {
		fmt.Println("Challenge 5 failed.")
		fmt.Println(err)
	}
	err = challenge7()
	if err != nil {
		fmt.Println("Challenge 7 failed.")
		fmt.Println(err)
	}
}
