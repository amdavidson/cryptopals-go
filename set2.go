package cryptopals

import "errors"

// Challenge 9
func pad(data []byte, padding byte, length int) ([]byte, error) {
	if len(data) >= length {
		return nil, errors.New("Padding input is longer than desired length")
	}

	for length > len(data) {
		data = append(data, padding)
	}

	return data, nil

}
