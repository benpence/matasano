package problems

import (
	"testing"

    "matasano/encoding"
    "matasano/bits"
    "matasano/crack/xorbyte"
)

const input3 string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
const output3 string = "Cooking MC's like a pound of bacon"

func Test3(t *testing.T) {
    // Decode ciphertext
    ciphertext, err := encoding.DecodeHex([]byte(input3))
    if err != nil {
        t.Error("Decoding", input3, "failed")
    }

    // Attempt to find byte most likely xor'd with ciphertext
    closestByte := xorbyte.FindByte(ciphertext, xorbyte.EnglishScorer)

    // Construct key from most likely byte
    key := make([]byte, len(ciphertext))
    for i := 0; i < len(key); i++ {
        key[i] = closestByte
    }

    // Decrypt ciphertext
    plaintext, _ := bits.Xor(ciphertext, key)

    // output plaintext
    if string(plaintext) != output3 {
        t.Error("Incorrect key chosen:", string(closestByte), string(plaintext))
    }
}
