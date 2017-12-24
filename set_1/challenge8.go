package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"
)

func challenge8() error {

	rawdata, err := ioutil.ReadFile("8.txt")
	if err != nil {
		return err
	}

	splits := strings.Split(string(rawdata), "\n")

	blocksize := 16

	encodedString := ""
	score := float64(len(splits[0]) / blocksize)

	for _, str := range splits {
		decodedBytes, err := hex.DecodeString(str)
		if err != nil {
			fmt.Println("Could not decode hex string")
			return err
		}

		blocks, err := chunk(decodedBytes, blocksize)
		if err != nil {
			fmt.Println("Could not chunk dataset")
			return err
		}

		temp := make(map[string]int)

		for i := 0; i < len(blocks); i++ {
			temp[string(blocks[i])] = 0
		}

		strScore := float64(len(temp)) / float64(len(blocks))

		if strScore < score {
			score = strScore
			encodedString = str
		}

	}

	fmt.Println("Encoded string is:", encodedString)
	fmt.Println("Unique block ratio was:", score)

	fmt.Println("Challenge 8 success!")
	return nil
}
