package main

import (
	"fmt"
)

func challenge1() {
	fmt.Println("Running Challenge 1")
	source := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	check := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	encoded, err := hextob64(source)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Check string: ", check)
	fmt.Println("Encoded string: ", encoded)

	if encoded == check {
		fmt.Println("Challenge 1 success!")
	} else {
		fmt.Println("Challenge 1 failed.")
	}

}
