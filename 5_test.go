package matasano

import (
	"testing"

	"matasano/bits"
	"matasano/encoding"
)

const input string = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
const key string = "ICE"
const output string = "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

func Test5(t *testing.T) {
    result := encoding.EncodeHex(bits.RepeatingXor([]byte(input), []byte(key)))
    if string(result) != output {
        t.Error(string(result), "!=", output)
	}
}
