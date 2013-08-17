package encoding

import (
	"testing"
)

var correctHexes map[string][]byte = map[string][]byte{
	"":                 []byte{},
	"0123456789abcdef": []byte{1, 35, 69, 103, 137, 171, 205, 239},
}

var invalidHexes []string = []string{
	"hijklmnopqrstuvwxyz",
	"abcdefg",
	"\n",
	"\t",
	" ",
}

func TestEncodeHex(t *testing.T) {
	for str, bytes := range correctHexes {
		result := EncodeHex(bytes)
		if string(result) != str {
			t.Error(string(result), "!=", str)
		}
	}
}

func TestDecodeHex(t *testing.T) {
	var (
		result []byte
		err    error
	)

	for str, bytes := range correctHexes {
		result, err = DecodeHex([]byte(str))

		if err != nil {
			t.Error("Error non-nil:", err)
			continue
		}

		if len(result) != len(bytes) {
			t.Error("Incorrect length")
			continue
		}

		for i, _ := range result {
			if result[i] != bytes[i] {
				t.Error(str, "decoded to", result)
				break
			}
		}
	}

	for _, invalidHex := range invalidHexes {
		if _, err = DecodeHex([]byte(invalidHex)); err == nil {
			t.Error(invalidHex, "did not return error")
		}
	}
}
