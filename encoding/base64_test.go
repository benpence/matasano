package encoding

import (
	"testing"

	"matasano/encoding"
)

var correctEncodings map[string][]byte = map[string][]byte{
	"": []byte{},
	"thequickbrownfoxjumpsoverthelazydogTHEQUICKBROWNFOXJUMPSOVERTHELAZYDOG0123456789+/th": []byte{
		182, 23, 170, 186, 39, 36, 110, 186,
		48, 157, 250, 49, 142, 233, 169, 178,
		139, 222, 174, 216, 94, 149, 172, 242,
		118, 136, 19, 28, 68, 20, 32, 34,
		129, 68, 229, 141, 20, 229, 201, 80,
		195, 210, 57, 81, 17, 76, 113, 11,
		1, 150, 3, 56, 109, 53, 219, 126,
		57, 235, 191, 61, 251, 251, 97},
}

var invalidEncodings []string = []string{
	"a",
	"ab",
	"abc",
	"!&*(",
}

func TestEncodeBase64(t *testing.T) {
	for str, bytes := range correctEncodings {
		result := encoding.EncodeBase64(bytes)
		if string(result) != str {
			t.Error(string(result), "!=", str)
		}
	}
}

func TestDecodeBase64(t *testing.T) {
	var (
		result []byte
		err    error
	)

	for str, bytes := range correctEncodings {
		result, err = encoding.DecodeBase64([]byte(str))

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

	for _, invalidEncoding := range invalidEncodings {
		if _, err = encoding.DecodeBase64([]byte(invalidEncoding)); err == nil {
			t.Error(invalidEncoding, "did not return error")
		}
	}
}
