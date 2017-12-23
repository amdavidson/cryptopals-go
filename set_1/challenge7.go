package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func challenge7() error {

	key := []byte("YELLOW SUBMARINE")

	bytes, err := ioutil.ReadFile("7.txt")
	if err != nil {
		return err
	}

	data, err := base64.StdEncoding.DecodeString(string(bytes))
	if err != nil {
		fmt.Println("File could not be base64 decoded")
	}

	decrypted, err := decryptECB(data, key)
	if err != nil {
		return err
	}

	fmt.Println(string(decrypted))

	fmt.Println("Challenge 7 success!")

	return nil
}
