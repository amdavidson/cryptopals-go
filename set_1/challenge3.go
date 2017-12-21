package cryptopals

import (
	"fmt"
)

func challenge3() error {
	fmt.Println("Running Challenge 3")
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	var err error
	strings := make([][]byte, 256)
	for i := 0x00; i <= 0xFF; i++ {
		strings[i] = make([]byte, len(input))
		strings[i], err = xorSingleChar(input, byte(i))
		if err != nil {
			return err
		}
	}

	score := 0
	char := ""
	out := ""

	for i := 0x00; i < 0xFF; i++ {
		iScore, err := countLowercase(string(strings[i]))
		if err != nil {
			fmt.Println("Could not count lowercase chars.")
			fmt.Println(err)
		}
		if score < iScore {
			score = iScore
			char = string(i)
			out = string(strings[i])
		}
	}

	fmt.Println("Highest score was", char, "with a score of", score)
	fmt.Println("String is:", out)

	fmt.Println("Challenge 3 success!")

	return nil
}
