package main

import (
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
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

// Challenge 7

func decryptECB(ct []byte, key []byte) ([]byte, error) {
	pt := make([]byte, len(ct))

	cipher, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key could not be used to make a new cipher")
		return nil, err
	}

	if len(ct)%len(key) != 0 {
		return nil, errors.New("cipher text is not a multiple of the key length")
	}

	blocksize := len(key)

	for blockstart, blockend := 0, blocksize; blockstart < len(ct); blockstart, blockend = blockstart+blocksize, blockend+blocksize {
		cipher.Decrypt(pt[blockstart:blockend], ct[blockstart:blockend])
	}

	return pt, nil
}

// Challenge 8

func chunk(data []byte, blocksize int) ([][]byte, error) {
	out := make([][]byte, 0)

	for i := 0; i < (len(data) / blocksize); i++ {
		out = append(out, data[(i*blocksize):((i+1)*blocksize)])
	}

	return out, nil

}


// Challenge 9

func pad(data []byte, padding byte, length int) []byte, error {
	if len(data) >= length {
		return []byte, errors.New("Padding input is longer than desired length")
	}

	for i := 0; i < (length - len(data)); i++ {
		data = append(data, padding)
	}

	if len(data) != len(padding) {
		return []byte, errors.New("Padding did not result in correct length")
	}

	return data, nil
	
}
