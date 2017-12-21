package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
    "errors"
)

// Set 1 Functions

// Challenge 1

func hextob64(message string) (string, error) {
	decoded, err := hex.DecodeString(message)
	if err != nil {
		return "", err
	}
	encoded := base64.StdEncoding.EncodeToString(decoded)
	return encoded, nil
}

// Challenge 2

func xorHex(inputstr string, keystr string) ([]byte, error) {
	input, err := hex.DecodeString(inputstr)
	if err != nil {
		fmt.Println("Failed to decode", inputstr)
		return make([]byte, 0), err
	}
	key, err := hex.DecodeString(keystr)
	if err != nil {
		fmt.Println("Failed to decode", keystr)
		return make([]byte, 0), err
	}
	out := make([]byte, len(input))
	for i, v := range input {
		out[i] = v ^ key[i]
	}

	return out, nil

}

//Challenge 3

func xorSingleChar(inputstr string, key byte) ([]byte, error) {
	input, err := hex.DecodeString(inputstr)
	if err != nil {
		fmt.Println("Failed to decode", inputstr)
		return make([]byte, 0), err
	}
	out := make([]byte, len(input))
	for i, v := range input {
		out[i] = v ^ key
	}

	return out, nil
}

func countLowercase(inputstr string) (int, error) {
	bytes := []byte(inputstr)
	count := 0
	for i := 0; i < len(bytes); i++ {
		if bytes[i] >= 'a' && bytes[i] <= 'z' {
			count++
		}
	}
	return count, nil
}

// Challenge 5

func xorRepeatingKey(inputString string, key []byte) ([]byte, error) {
	inputBytes := []byte(inputString)

	outputBytes := make([]byte, len(inputString))

	for i, inputByte := range inputBytes {
		xorIndex := i % 3
		outputBytes[i] = inputByte ^ key[xorIndex]
	}

	return outputBytes, nil
}

// Challenge 6

func hammingDistance(string1 []byte, string2 []byte) (int, error) {
	distance := 0

    if len(string1) != len(string2) {
        return 0, errors.New("String lengths are unmatched")
    }

	for i, val1 := range string1 {
		if val1 == string2[i] {
			continue
		} else {
			distance = distance + 1
		}
	}

	return distance, nil
}
