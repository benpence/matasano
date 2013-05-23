package bits

import (
	"errors"
)

func Xor(first, second []byte) ([]byte, error) {
	if len(first) != len(second) {
		return nil, errors.New("Slices are different lengths")
	}

	output := make([]byte, len(first))

	for i := 0; i < len(first); i++ {
		output[i] = first[i] ^ second[i]
	}

	return output, nil
}
