package main

import (
	"errors"
	"fmt"
    "encoding/hex"
)

func challenge5() error {

	inputString := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`

	xorKey := "ICE"

	outputCheck := `0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`

	outputBytes, err := xorRepeatingKey(inputString, []byte(xorKey))
    outputString := hex.EncodeToString(outputBytes)

	if err != nil {
		fmt.Println("Repeating key XOR function failed")
		return err
	}

	if outputString != outputCheck {
		fmt.Println("Expected:", outputCheck)
		fmt.Println("Got:", outputString)
		return errors.New("XOR function did not match checkstring")
	}

    fmt.Println("Input String:",inputString)
    fmt.Println("XOR key:",xorKey)
    fmt.Println("Check string:",outputCheck)
    fmt.Println("Computed string:",outputString)
	fmt.Println("Challenge 5 Success!")
	return nil

}
