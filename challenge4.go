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

    for _, testCase := range testCases {
        for i := 0; i <= 255; i++ {
			testString, err := xorSingleChar(testCase, byte(i))
			if err != nil {
				fmt.Println("XOR failed")
				return err
			}
			testScore, err := countLowercase(string(testString))
			if err != nil {
				fmt.Println("Scoring failed")
				return err
			}
			if score < testScore {
				score = testScore
				char = string(i)
				out = string(testString)
			}
		}
	}

	fmt.Println("High scoring string is:", strings.Trim(out, "\n"))
	fmt.Println("Score is:", score)
	fmt.Println("XOR char is:", char)

	fmt.Println("Challenge 4 success!")
	return nil
}
