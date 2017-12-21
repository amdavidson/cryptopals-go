package main

import (
	"errors"
	"fmt"
)

func challenge1() error {
	fmt.Println("Running Challenge 1")
	source := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	check := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	encoded, err := hextob64(source)
	if err != nil {
		return err
	}

	fmt.Println("Check string: ", check)
	fmt.Println("Encoded string: ", encoded)

	if encoded == check {
		fmt.Println("Challenge 1 success!")
	} else {
		return errors.New("Encoded does not equal check value")
	}

	return nil
}
