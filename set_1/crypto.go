package main

import (
	"encoding/base64"
	"encoding/hex"
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

func xorhex(inputstr string, keystr string) ([]byte, error) {
	input, err := hex.DecodeString(inputstr)
	if err != nil {
		return make([]byte, 0), err
	}
	key, err := hex.DecodeString(keystr)
	if err != nil {
		return make([]byte, 0), err
	}
	out := make([]byte, len(input))
	for i, v := range input {
		out[i] = v ^ key[i]
	}

	return out, nil

}
