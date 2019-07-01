package cryptopals

import (
	"bytes"
	"testing"
)

func TestProblem9(t *testing.T) {
	input := []byte("YELLOW SUBMARINE")
	padding := byte(0x04)
	check := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")

	output, err := pad(input, padding, 20)
	if err != nil {
		t.Fatal(err)
	}

	if len(output) != len(check) {
		t.Error("Lengths do not match", len(output), len(check))
	}

	if !bytes.Equal(output, []byte("YELLOW SUBMARINE\x04\x04\x04\x04")) {
		t.Error("Did not pad properly", output)
	}
}
