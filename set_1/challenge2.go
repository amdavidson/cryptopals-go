package main

import (
	"encoding/hex"
	"fmt"
)

func challenge2() {
	fmt.Println("Running Challenge 2.")
	string1 := "1c0111001f010100061a024b53535009181c"
	string2 := "686974207468652062756c6c277320657965"
	check := "746865206b696420646f6e277420706c6179"
	fmt.Println("Check string: ", check)
	xor, err := xorhex(string1, string2)
	if err != nil {
		fmt.Println("Challenge 2 failed")
		return
	}
	out := hex.EncodeToString(xor)

	fmt.Println("Result: ", out)

	if out == check {
		fmt.Println("Challenge 2 Success!")
	} else {
		fmt.Println("Challenge 2 failed.")
	}

}
