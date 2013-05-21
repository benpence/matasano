package matasano

import (
	"testing"

	"matasano/bits"
	"matasano/encoding"
)

const input5 string = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
const key5 string = "ICE"
const output5 string = "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

func Test5(t *testing.T) {
    ciphertext, err := bits.RepeatingXor([]byte(input5), []byte(key5))
    if err != nil {
        t.Error("bits.RepeatingXor returned error:", err)
    }

    result := encoding.EncodeHex(ciphertext)
    if string(result) != output5 {
        t.Error(string(result), "!=", output5)
	}
}
