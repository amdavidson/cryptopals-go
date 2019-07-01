package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"strings"
	"testing"
)

func TestProblem1(t *testing.T) {
	res, err := hextob64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if err != nil {
		t.Fatal(err)
	}

	if res != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
		t.Error("Wrong output string", res)
	}
}

func TestProblem2(t *testing.T) {
	string1 := "1c0111001f010100061a024b53535009181c"
	string2 := "686974207468652062756c6c277320657965"
	check := "746865206b696420646f6e277420706c6179"

	xor, err := xorHex(string1, string2)
	if err != nil {
		t.Fatal(err)
	}

	out := hex.EncodeToString(xor)

	if out != check {
		t.Error("String does not match checkvalue", err)
	}
}

func TestProblem3(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	var err error
	// build series of character strings for xor
	strings := make([][]byte, 256)
	for i := 0x00; i <= 0xFF; i++ {
		strings[i] = make([]byte, len(input))
		strings[i], err = xorSingleChar(input, byte(i))
		if err != nil {
			t.Fatal(err)
		}
	}

	score := 0
	char := ""
	out := ""

	for i := 0x00; i < 0xFF; i++ {
		iScore, err := countLowercase(string(strings[i]))
		if err != nil {
			t.Fatal(err)
		}
		if score < iScore {
			score = iScore
			char = string(i)
			out = string(strings[i])
		}
	}

	t.Logf("%s: %s", char, out)
}

func TestProblem4(t *testing.T) {
	data, err := ioutil.ReadFile("4.txt")
	if err != nil {
		t.Fatal(err)
	}
	testCases := strings.Split(string(data), "\n")

	score := 0
	char := ""
	out := ""

	for _, testCase := range testCases {
		for i := 0; i <= 255; i++ {
			testString, err := xorSingleChar(testCase, byte(i))
			if err != nil {
				t.Fatal(err)
			}
			testScore, err := countLowercase(string(testString))
			if err != nil {
				t.Fatal(err)
			}
			if score < testScore {
				score = testScore
				char = string(i)
				out = string(testString)
			}
		}
	}

	t.Logf("%s: %s", char, out)
}

func TestProblem5(t *testing.T) {
	inputString := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`

	xorKey := "ICE"

	outputCheck := `0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`

	outputBytes, err := xorRepeatingKey(inputString, []byte(xorKey))
	outputString := hex.EncodeToString(outputBytes)

	if err != nil {
		t.Fatal(err)
	}

	if outputString != outputCheck {
		t.Error("Did not match checkvalue", outputString)
	}
}

func TestProblem6(t *testing.T) {
	t.Log("WARN: not implemented")
}

func TestProblem7(t *testing.T) {
	key := []byte("YELLOW SUBMARINE")

	bytes, err := ioutil.ReadFile("7.txt")
	if err != nil {
		t.Fatal(err)
	}

	data, err := base64.StdEncoding.DecodeString(string(bytes))
	if err != nil {
		t.Fatal(err)
	}

	decrypted, err := decryptECB(data, key)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", string(decrypted))
}

func TestProblem8(t *testing.T) {
	rawdata, err := ioutil.ReadFile("8.txt")
	if err != nil {
		t.Fatal(err)
	}

	splits := strings.Split(string(rawdata), "\n")

	blocksize := 16

	encodedString := ""
	score := float64(len(splits[0]) / blocksize)

	for _, str := range splits {
		decodedBytes, err := hex.DecodeString(str)
		if err != nil {
			t.Fatal(err)
		}

		blocks, err := chunk(decodedBytes, blocksize)
		if err != nil {
			t.Fatal(err)
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

	t.Logf("%s", encodedString)
	t.Logf("%f", score)
}
