package matasano

import (
	"testing"

	"matasano/encoding"
)

const input string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
const output string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

func Test1(t *testing.T) {
    raw, err := encoding.DecodeHex([]byte(input))
    if err != nil {
        t.Error("Decoding", input, "failed")
    }

    if encoded := encoding.EncodeBase64(raw); string(encoded) != output {
        t.Error(string(encoded), "!=", output)
	}
}
