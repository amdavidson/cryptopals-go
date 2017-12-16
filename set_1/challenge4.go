package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func challenge4() error {
	data, err := ioutil.ReadFile("4.txt")
	if err != nil {
		return err
	}
	testCases := strings.Split(string(data), "\n")

	score := 0
	char := ""
	out := ""

	for j := 0; j < len(testCases); j++ {

		for i := 0x00; i < 0xFF; i++ {
			ijString, err := xorSingleChar(testCases[j], byte(i))
			if err != nil {
				fmt.Println("XOR failed")
				return err
			}
			ijScore, err := countLowercase(string(ijString))
			if err != nil {
				fmt.Println("Scoring failed")
				return err
			}
			if score < ijScore {
				score = ijScore
				char = string(i)
				out = string(ijString)
			}
		}
	}

	fmt.Println("High scoring string is:", strings.Trim(out, "\n"))
	fmt.Println("Score is:", score)
	fmt.Println("XOR char is:", char)

	fmt.Println("Challenge 4 success!")
	return nil
}
