package matasano

import (
	"testing"

    "matasano/encoding"
    "matasano/bits"
    "matasano/crack"
)

const input string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
var englishFrequencyDistribution []float32 = []float32 {
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0.08167, 0.01492, 0.02782, 0.04253, 0.12702, 0.02228, 0.02015, 0.06094, 0.06966, 0.00153, 0.00772, 0.04025, 0.02406, 0.06749, 0.07507,
    0.01929, 0.00095, 0.05987, 0.06327, 0.09056, 0.02758, 0.00978, 0.02360, 0.00150, 0.01974, 0.00074, 0, 0, 0, 0, 0,
    0, 0.08167, 0.01492, 0.02782, 0.04253, 0.12702, 0.02228, 0.02015, 0.06094, 0.06966, 0.00153, 0.00772, 0.04025, 0.02406, 0.06749, 0.07507,
    0.01929, 0.00095, 0.05987, 0.06327, 0.09056, 0.02758, 0.00978, 0.02360, 0.00150, 0.01974, 0.00074, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 }

const output string = "Cooking MC's like a pound of bacon"

func Test3(t *testing.T) {
    // Decode ciphertext
    ciphertext, err := encoding.DecodeHex([]byte(input))
    if err != nil {
        t.Error("Decoding", input, "failed")
    }

    // Attempt to find byte most likely xor'd with ciphertext
    closestByte, err := crack.XorByte(ciphertext, englishFrequencyDistribution)
    if err != nil {
        t.Error("Failed to find most common letter:", err)
    }

    // Construct key from most likely byte
    key := make([]byte, len(ciphertext))
    for i := 0; i < len(key); i++ {
        key[i] = closestByte
    }

    // Decrypt ciphertext
    plaintext, _ := bits.Xor(ciphertext, key)

    // Output plaintext
    if string(plaintext) != output {
        t.Error("Incorrect key chosen:", string(closestByte), string(plaintext))
    }
}
